<script>
    import {getContext} from 'svelte';
    import ConfirmModal from './ConfirmModal.svelte';
    import SettingsModal from './SettingsModal.svelte';
    import FilterIcon from '../icons/filter.svelte';
    import AddIcon from '../icons/add.svelte';
    import SettingsIcon from '../icons/settings.svelte';
    import QuitIcon from '../icons/quit.svelte';
    import {toggleFilterPane} from "../stores";
    import {addResource} from "../stores";

    const {open} = getContext('simple-modal');
    const showQuitModal = () => {
        open(ConfirmModal, {message: "Do you really want to quit?", onOkay: () => {
            window.backend.quit()
        }});
    };
    const showSettingsModal = () => {
        open(SettingsModal, {});
    }

</script>

<style>
    .menu-section {
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: 16px;
        /*border: var(--border);*/
    }

    .menu-icon {
        display: inline-flex;
        border: var(--border);
        border-radius: var(--border-radius);
        box-shadow: 1px 1px 5px var(--bg-color-dark);
        padding: 6px;
        cursor: pointer;
    }

    .menu-icon:hover {
        background: var(--bg-color2);
    }

    .menu-icon:active {
        background: var(--border-color);
        box-shadow: none;
    }

    nav {
        height: var(--nav-height);
        box-sizing: border-box;
        /*border-bottom: var(--border);*/
        box-shadow: 0 1px 5px var(--bg-color-dark);

        padding: var(--padding);
        /*flex: 0 0 auto;*/
        display: flex;
        flex-flow: row;
        align-items: center;
        justify-content: space-between;

        /* disable text selection: */
        user-select: none;
        -webkit-user-select: none;
    }
</style>

<nav>
    <div class="menu-section">
        <div class="menu-icon" on:click={() => {{
            toggleFilterPane();
        }}}>
            <FilterIcon/>
        </div>
        <div class="menu-icon" on:click={() => addResource()}>
            <AddIcon/>
        </div>
    </div>
    <div class="menu-section">
        <div class="app-name">
            <h2>OsintDB</h2>
        </div>
    </div>
    <div class="menu-section">
        <div class="menu-icon" on:click={showSettingsModal}>
            <SettingsIcon/>
        </div>
        <div class="menu-icon" on:click={showQuitModal}>
            <QuitIcon/>
        </div>
    </div>
</nav>
