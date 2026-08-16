package services

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"portfolio-os/internal/models"
)

type AnalyticsService struct {
	mu     sync.RWMutex
	visits map[string]*models.Visit
}
type AnalyticsStats struct {
	TotalVisits         int             `json:"total_visits"`
	UniqueVisitors      int             `json:"unique_visitors"`
	AverageDuration     time.Duration   `json:"-"`
	AverageDurationText string          `json:"average_duration"`
	DeviceBreakdown     map[string]int  `json:"device_breakdown"`
	PageViews           map[string]int  `json:"page_views"`
	RecentVisits        []*models.Visit `json:"recent_visits"`
}

func NewAnalyticsService() *AnalyticsService {
	return &AnalyticsService{
		visits: make(map[string]*models.Visit),
	}
}

func (s *AnalyticsService) StartVisit(visit *models.Visit) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.visits[visit.ID] = visit
}

func (s *AnalyticsService) UpdateDuration(
	visitID string,
	duration time.Duration,
) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	visit, exists := s.visits[visitID]

	if !exists {
		return false
	}

	visit.Duration = duration

	return true
}

func (s *AnalyticsService) GetVisits() []*models.Visit {
	s.mu.RLock()
	defer s.mu.RUnlock()

	visits := make([]*models.Visit, 0, len(s.visits))

	for _, visit := range s.visits {
		visits = append(visits, visit)
	}

	return visits
}

func (s *AnalyticsService) GetStats() *AnalyticsStats {
	s.mu.RLock()
	defer s.mu.RUnlock()

	stats := &AnalyticsStats{
		TotalVisits:     len(s.visits),
		DeviceBreakdown: make(map[string]int),
		PageViews:       make(map[string]int),
	}

	uniqueVisitors := make(map[string]struct{})

	var totalDuration time.Duration

	visits := make([]*models.Visit, 0, len(s.visits))

	for _, visit := range s.visits {
		visits = append(visits, visit)

		// Unique visitor based on anonymized IP.
		uniqueVisitors[visit.IPAddress] = struct{}{}

		// Device statistics.
		stats.DeviceBreakdown[visit.DeviceCategory]++

		// Page statistics.
		stats.PageViews[visit.Page]++

		// Duration.
		totalDuration += visit.Duration
	}

	stats.UniqueVisitors = len(uniqueVisitors)

	if stats.TotalVisits > 0 {
		stats.AverageDuration = totalDuration / time.Duration(stats.TotalVisits)
		stats.AverageDurationText = FormatDuration(stats.AverageDuration)
	}

	// Most recent visits first.
	sort.Slice(visits, func(i, j int) bool {
		return visits[i].Timestamp.After(visits[j].Timestamp)
	})

	// Return only the latest 10 visits.
	if len(visits) > 10 {
		visits = visits[:10]
	}

	stats.RecentVisits = visits

	return stats
}

func FormatDuration(duration time.Duration) string {
	totalSeconds := int(duration.Seconds())

	if totalSeconds < 60 {
		return fmt.Sprintf("%d sec", totalSeconds)
	}

	minutes := totalSeconds / 60
	seconds := totalSeconds % 60

	if minutes < 60 {
		if seconds == 0 {
			return fmt.Sprintf("%d min", minutes)
		}

		return fmt.Sprintf("%d min %d sec", minutes, seconds)
	}

	hours := minutes / 60
	minutes %= 60

	if minutes == 0 {
		return fmt.Sprintf("%d hr", hours)
	}

	return fmt.Sprintf("%d hr %d min", hours, minutes)
}
