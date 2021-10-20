<script>
    // footer for status stuff, notifications and co
    // idea: showing X of Y tools (x: number of filtered tools, y: number of all tools)
    // idea: "dropdown" menu for lists (all, links, extensions, apps, custom lists)
    // idea: either show notifications/errors or button to notification/error log (which flashes red when new message arrived)
    import Notifications from "./notifications/Notifications.svelte";
    import {statsStore} from "../stores";

    let stats;
    statsStore.subscribe(s => {
        stats = s;
    })
</script>

<style>
    .footer {
        height: 20px;
        box-sizing: border-box;
        border-top: 1px solid var(--border-color);
        box-shadow: 0 -1px 5px var(--bg-color-dark);
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        padding-left: 12px;
        padding-right: 12px;

        /* disable text selection: */
        user-select: none;
        -webkit-user-select: none;
    }
</style>

<footer>
    <div class="footer">
        <div>{stats.filtersCount} active {stats.filtersCount === 1 ? "filter" : "filters"}</div>
        <div>showing {stats.filteredCount} of {stats.resourcesCount}</div>
        <Notifications/>
    </div>
</footer>