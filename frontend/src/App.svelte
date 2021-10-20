<script>
    // TODO: style modals with CSS to match theme
    // TODO: status information in footer
    // idea: add keybindings
    //		"/": for searching with term
    //   	"Ctrl+F", "Alt+1": for opening filters
    // 		"Ctrl+A": for adding ressource [follow up shortcuts for type to add]
    //      "Ctrl+Alt+S": for opening settings
    //      "Ctrl+Q": for quitting (maybe with confirm dialog)
    //      "Esc": for closing popups/modals
    import Modal from 'svelte-simple-modal';
    import TopBar from './components/TopBar.svelte';
    import Footer from './components/Footer.svelte';
    import FilterPane from "./components/FilterPane.svelte";
    import TablePane from "./components/TablePane.svelte";

    import {initStores} from "./stores";

    initStores();

    const modalWindowStyle = {
        width: "fit-content",
        padding: "10px",
    };

</script>

<style>
    :root {
        --bg-color: #2e3440;
        --bg-color2: #434c5e;
        --bg-color3: #4c566a;
        --bg-color-light-dark: #2a2f3c;
        --bg-color-dark: #1d2330;
        --bg-inverse-color: #eceffa;
        --bg-inverse-color2: #e5e9f0;
        --bg-inverse-color3: #d8dee9;
        --bg-input-color: #242933;
        --border-color: #3b4252;
        --text-color: #eceff4;
        --text-color2: #d8dee9;
        --text-color3: #bcc7d9;
        --text-inverse-color: #2e3440;
        --primary-color: #88c0d0;
        --accent-color: #8fbcbb;
        --accent-color2: #81a1c1;
        --accent-color3: #5e81ac;
        --red-color: #bf616a;
        --orange-color: #d08770;
        --yellow-color: #ebcb8b;
        --green-color: #a3be8c;
        --purple-color: #b48ead;
        --padding: 12px;
        --border: 1px solid #3b4252;
        --border-radius: 3px;
        --font-size: 10pt;
        --footer-height: 20px;
        --nav-height: 54px;
        --panes-top-margin: 8px;
    }

    :global(html,body) {
        margin: 0;
        height: 100%;
        font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto", "Oxygen",
        "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue",
        sans-serif;
        font-size: var(--font-size);
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        /*user-select: none;*/
        /*-webkit-user-select: none;*/
        background-color: var(--bg-color);
        color: var(--text-color);
        cursor: default;
        /*-ms-overflow-style: -ms-autohiding-scrollbar;*/
    }

    .app {
        height: 100%;
        display: flex;
        flex-flow: column;
    }

    .nav {
        height: var(--nav-height);
    }

    .panes {
        height: calc(100% - var(--nav-height) - var(--panes-top-margin) - var(--footer-height));
        display: flex;
        flex-flow: row;
        margin-top: 8px;
    }
</style>

<main class="app">
    <Modal styleWindow="{modalWindowStyle}">
        <div class="nav">
            <TopBar/>
        </div>
        <div class="panes">
            <FilterPane/>
            <TablePane/>
        </div>
    </Modal>
    <Footer/>
</main>
