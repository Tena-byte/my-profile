document.addEventListener("DOMContentLoaded", () => {
    const clientView = document.getElementById("client-view");
    const developerView = document.getElementById("developer-view");

    const toggleButton = document.getElementById("view-toggle");
    const devViewToggle = document.getElementById("dev-view-toggle");
    const viewLabel = document.getElementById("view-label");

    const editorTabs = document.querySelectorAll(".editor-tab");
    const fileItems = document.querySelectorAll(".file-item");
    const codePanels = document.querySelectorAll(".code-panel");

    /*
     * --------------------------------------------------
     * VIEW TOGGLE
     * --------------------------------------------------
     */

    function setView(view) {
    const isDeveloper = view === "developer";

    const activeView = isDeveloper
        ? developerView
        : clientView;

    const inactiveView = isDeveloper
        ? clientView
        : developerView;

    inactiveView.classList.remove("view-visible");
    inactiveView.classList.add("view-hidden");

    setTimeout(() => {
        inactiveView.hidden = true;

        activeView.hidden = false;

        requestAnimationFrame(() => {
            activeView.classList.remove("view-hidden");
            activeView.classList.add("view-visible");
        });
    }, 150);

    if (viewLabel) {
        viewLabel.textContent = isDeveloper
            ? "Developer View"
            : "Client View";
    }

    if (devViewToggle) {
        devViewToggle.textContent = isDeveloper
            ? "Client View"
            : "Developer View";
    }

    localStorage.setItem("portfolio-view", view);
}


    /*
     * --------------------------------------------------
     * CLIENT VIEW TO DEVELOPER VIEW
     * --------------------------------------------------
     */

    if (toggleButton) {
        toggleButton.addEventListener("click", () => {
            const currentView =
                localStorage.getItem("portfolio-view") || "client";

            setView(
                currentView === "client"
                    ? "developer"
                    : "client"
            );
        });
    }


    /*
     * --------------------------------------------------
     * DEVELOPER VIEW TO CLIENT VIEW
     * --------------------------------------------------
     */

    if (devViewToggle) {
        devViewToggle.addEventListener("click", () => {
            const currentView =
                localStorage.getItem("portfolio-view") || "client";

            setView(
                currentView === "client"
                    ? "developer"
                    : "client"
            );
        });
    }


    /*
     * --------------------------------------------------
     * EDITOR TAB SWITCHING
     * --------------------------------------------------
     */

    function openTab(tabName) {

        /*
         * Update editor tabs
         */

        editorTabs.forEach((tab) => {
            tab.classList.toggle(
                "active",
                tab.dataset.tab === tabName
            );
        });


        /*
         * Update explorer files
         */

        fileItems.forEach((file) => {
            file.classList.toggle(
                "active",
                file.dataset.tab === tabName
            );
        });


        /*
         * Update editor panels
         */

        codePanels.forEach((panel) => {
            panel.classList.toggle(
                "active",
                panel.dataset.panel === tabName
            );
        });


        /*
         * Remember currently open file
         */

        localStorage.setItem(
            "portfolio-open-file",
            tabName
        );


        /*
         * Update status bar
         */

        updateStatusBar(tabName);
    }


    /*
     * --------------------------------------------------
     * TAB CLICK EVENTS
     * --------------------------------------------------
     */

    editorTabs.forEach((tab) => {
        tab.addEventListener("click", () => {
            openTab(tab.dataset.tab);
        });
    });


    /*
     * --------------------------------------------------
     * EXPLORER CLICK EVENTS
     * --------------------------------------------------
     */

    fileItems.forEach((file) => {
        file.addEventListener("click", () => {
            openTab(file.dataset.tab);
        });
    });


    /*
     * --------------------------------------------------
     * STATUS BAR
     * --------------------------------------------------
     */

    function updateStatusBar(tabName) {
        const statusFile =
            document.querySelector(".status-file");

        if (!statusFile) {
            return;
        }

        const fileNames = {
            home: "Home.go",
            about: "AboutMe.go",
            skills: "Skills.go",
            services: "Services.go",
            education: "Education.go",
            certificates: "Certificates.go",
            projects: "Projects.go",
            gallery: "Gallery.go",
            contact: "Contact.go"
        };

        statusFile.textContent =
            fileNames[tabName] || "Go";
    }


    /*
     * --------------------------------------------------
     * DYNAMIC LINE NUMBERS
     * --------------------------------------------------
     */

    function updateLineNumbers() {
        const codeEditors =
            document.querySelectorAll(".code-editor");

        codeEditors.forEach((editor) => {
            const code =
                editor.querySelector("code");

            const lineNumbers =
                editor.querySelector(".line-numbers");

            if (!code || !lineNumbers) {
                return;
            }

            const lines =
                code.innerText.split("\n");

            lineNumbers.innerHTML = "";

            lines.forEach((_, index) => {
                const line =
                    document.createElement("span");

                line.textContent = index + 1;

                lineNumbers.appendChild(line);
            });
        });
    }


    /*
     * --------------------------------------------------
     * INITIAL STATE
     * --------------------------------------------------
     */

    const savedView =
        localStorage.getItem("portfolio-view") || "client";

    setView(savedView);


    const savedFile =
        localStorage.getItem("portfolio-open-file") || "home";

    openTab(savedFile);


    updateLineNumbers();





    /*
 * --------------------------------------------------
 * MOBILE EXPLORER
 * --------------------------------------------------
 */

const mobileExplorerToggle =
    document.getElementById("mobile-explorer-toggle");

const devSidebar =
    document.querySelector(".dev-sidebar");

if (mobileExplorerToggle && devSidebar) {

    mobileExplorerToggle.addEventListener("click", () => {

        const isOpen =
            devSidebar.classList.toggle("mobile-open");

        mobileExplorerToggle.setAttribute(
            "aria-expanded",
            isOpen
        );
    });


    fileItems.forEach((file) => {

        file.addEventListener("click", () => {

            if (window.innerWidth <= 768) {
                devSidebar.classList.remove("mobile-open");

                mobileExplorerToggle.setAttribute(
                    "aria-expanded",
                    "false"
                );
            }

          });

        });

    }
    
});