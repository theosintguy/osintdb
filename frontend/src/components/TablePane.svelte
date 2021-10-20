<script>
    // TODO: add options to enable/disable table cells (have second dropzone field with column names on "edit button" and everything in zone1/table is displayed and everything in zone2 is not)
    // TODO: check that mandatory fields are filled out on modification or addition of resource
    // TODO: improve way to add or modify enum-style fields (tags, browsers, licenses, type, regions)
    // TODO: for Applications: check if is installed? if it is: start it by clicking the title or a special icon

    import {filtersStore, resourcesStore, statsStore, subscribeAddResources} from "../stores";
    import {buildFlexIndex, filter, sortColumns} from "../filtering";
    import {dndzone} from "svelte-dnd-action";
    import Wails from "@wailsapp/runtime";
    import {getContext} from "svelte";
    import ConfirmModal from "./ConfirmModal.svelte";

    let resources = [];
    let filtered;
    let filters = new Map();
    filters.set("title", "");
    filters.set("description", "");
    filters.set("url", "");
    filters.set("comment", "");
    filters.set("resource_type", []);
    filters.set("licenses", []);
    filters.set("browsers", []);
    filters.set("tags", []);
    filters.set("regions", []);

    let columnsAvailable = [
        {"id": 0, "active": false, "key": "id", "name": "ID"},
        {"id": 1, "active": true, "key": "title", "name": "Title"},
        {"id": 2, "active": true, "key": "resource_type", "name": "Type"},
        {"id": 3, "active": true, "key": "description", "name": "Description"},
        {"id": 4, "active": true, "key": "url", "name": "URL"},
        {"id": 5, "active": false, "key": "comment", "name": "Comment"},
        {"id": 6, "active": false, "key": "licenses", "name": "Licenses"},
        {"id": 7, "active": false, "key": "browsers", "name": "Browsers"},
        {"id": 8, "active": false, "key": "tags", "name": "Tags"},
        {"id": 9, "active": false, "key": "regions", "name": "Regions"},
    ];
    let columns;

    $: columns = columnsAvailable.filter(c => c.active);

    resourcesStore.subscribe(async jsonData => {
        resources = JSON.parse(jsonData);
        buildFlexIndex(resources);
        filtered = filter(resources, filters);
        updateStats();
    });

    function updateStats() {
        statsStore.update(stats => {
            stats.resourcesCount = resources.length;
            stats.filteredCount = filtered.length;
            return stats
        })
    }

    // window.backend.DB.UpdateAll(); // hack to trigger `resourcesStore` update

    filtersStore.subscribe(async f => {
        filters = f;
        filtered = filter(resources, filters);
    })

    // wrapper for sorting when clicking on column header
    function sort(column) {
        filtered = sortColumns(filtered, column);
    }

    function viewDetails(d) {
        details.data = d;
        details.active = true;
    }

    async function showAddResource() {
        if (details.active && details.editable) {
            cancelEditDetails()
        }
        let jsonData = await window.backend.DB.GetResourceTemplate()
        details.data = JSON.parse(jsonData)

        details.new = true;
        details.editable = true;
        details.active = true;

    }

    let resourceDetails;


    // dndzone stuff

    function handleDndConsider(e) {
        columns = e.detail.items;
    }

    function handleDndFinalize(e) {
        columns = e.detail.items;
    }

    async function reload() {
        await window.backend.DB.UpdateAll();
    }

    let mode = "table";

    let details = {"active": false, "data": undefined, "editable": false, "new": false};


    const {open} = getContext('simple-modal');
    const confirmDelete = async id => {
        open(ConfirmModal, {message: "Do you really want to delete the Resource?",
            onOkay: () => {
                window.backend.DB.DeleteResource(id);
                goBackFromDetails();
            }
        })
    }

    async function editDetails() {
        // TODO: find better way to edit: type, tags and "additionals" ["enum"-style with capability to add new enum types (apart from "type")] (with autocomplete for existing enum types?)
        details.editable = true;
    }

    async function cancelEditDetails() {
        details.editable = false;
        details.new = false;
    }

    async function submitEditDetails() {
        for (let field of Object.keys(details.data).filter(k => k !== "id")) {
            let id = `details-${field}`;
            let elem = document.getElementById(id);
            if (typeof details.data[field] === "string") {
                details.data[field] = elem.innerText;
            } else if (Array.isArray(details.data[field])) {
                details.data[field] = elem.innerText.split(",").map(item => item.trim()).filter(item => item !== "");
            }
        }
        // TODO: check that mandatory fields are filled out. (title, type) OR WHAT?
        let jsonData = JSON.stringify(details.data);
        if (details.new) {
            // add new resource to backend DB
            await window.backend.DB.AddResource(jsonData);
            details.new = false;
        } else {
            // modify existing resource in backend DB
            await window.backend.DB.ModifyResource(jsonData);
        }
        details.editable = false;
    }

    async function goBackFromDetails() {
        if (details.editable) {
            await cancelEditDetails();
        }
        details.new = false;
        details.active = false;
    }

    function fieldToName(field) {
        switch (field) {
            case "resource_type":
                return "Type"
            case "url":
                return "URL"
            default:
                return field.charAt(0).toUpperCase() + field.slice(1);
        }
    }

    function formatField(field) {
        if (Array.isArray(field)) {
            return field.join(", ");
        }
        return field;
    }

    subscribeAddResources(val => {
        if (val) { // hack to prevent trying to add resource on first value update
            showAddResource();
        }
    })
</script>

<style>
    .container {
        width: 100%;
        overflow-y: scroll;
        -ms-overflow-style: none;
        scrollbar-width: none;
        display: flex;
        flex-direction: column;
    }

    .container::-webkit-scrollbar {
        display: none;
    }

    table {
        border-collapse: collapse;
        overflow: hidden;
        /*box-shadow: 0 0 20px rgba(0,0,0,0.1);*/
        width: 100%;
        border-bottom: 1px solid var(--bg-color-light-dark);
        /*flex-grow: 1;*/
        position: relative;
        top: 0;
        bottom: 100px;
    }

    th, td {
        padding: 15px;
        /*max-width: 200px;*/
        white-space: pre-wrap;
    }

    th {
        text-align: left;
    }

    tr:not(thead tr):hover, tr:nth-child(even):hover {
        background-color: var(--bg-color2);
    }

    tr:nth-child(even), thead {
        background-color: var(--bg-color-light-dark);
    }

    thead th:hover {
        background-color: var(--bg-color2);
    }

    .not-found {
        margin: 20px;
    }

    .details-menu {
        display: flex;
        flex-direction: row;
        gap: 12px;
    }

    .details-btn {
        /*font-size: 50px;*/
        border: var(--border);
        border-radius: var(--border-radius);
        box-shadow: 1px 1px 5px var(--bg-color-dark);
        background: var(--bg-color);
        color: var(--text-color);
        padding: 5px 14px;
        text-align: center;
        cursor: pointer;

        /* disable text selection: */
        user-select: none;
        -webkit-user-select: none;
    }

    .details-btn:hover {
        background: var(--bg-color2);
    }

    .details-btn:active {
        background: var(--border-color);
        box-shadow: none;
    }

    .details-view {
        display: flex;
        flex-direction: column;
        margin-left: 12px;
        margin-right: 12px;
    }

    .details-infos {
        margin-top: 20px;
        display: flex;
        flex-direction: column;
    }

    .details-infos td {
        position: relative;
    }

    .details-infos tr td:first-child {
        width: 100px;
    }

    a {
        color: var(--text-color);
    }

    a:hover {
        color: var(--primary-color);
    }

</style>

<div class="container">
    {#if !details.active}
        {#if (typeof filtered !== 'undefined' && filtered.length > 0)}
            <table>
                <thead>
                <tr use:dndzone={{items: columns}} on:consider={handleDndConsider} on:finalize={handleDndFinalize}>
                    {#each columns as column(column.id)}
                        <th on:click={() => sort(column["key"])}>{column["name"]}</th>
                    {/each}
                </tr>
                </thead>
                <tbody>
                {#each Object.values(filtered) as entry(entry.id)}
                    <tr on:click={e => {{
                    if (e.target.tagName !== "A") {
                        // make links clickable without opening details view
                        viewDetails(entry);
                    }
                }}}>
                        {#each columns as column(column.id)}
                            {#if column["key"] === "url"}
                                <td><a href="#" on:click={Wails.Browser.OpenURL(entry["url"])}>{entry["url"]}</a></td>
                            {:else}
                                <td>{entry[column["key"]]}</td>
                            {/if}
                        {/each}
                    </tr>
                {/each}
                </tbody>
            </table>
        {:else}
            <div class="not-found" use:reload>
                <h2>No Resources found!</h2>
                <p>
                    Try to adjust your filters.
                </p>
            </div>
        {/if}
    {:else} <!-- details are active -->
        <div class="details-view">
            <div class="details-menu">
                <button class="details-btn back-btn" on:click={() => goBackFromDetails()}>&lt; Back</button>
                {#if !details.editable}
                    <button class="details-btn modify-btn" on:click={() => editDetails()}>Modify</button>
                {:else}
                    <button class="details-btn confirm-btn" on:click={() => submitEditDetails()}>
                        Confirm Changes
                    </button>
                    {#if !details.new}
                    <button class="details-btn cancel-btn" on:click={() => cancelEditDetails()}>
                        Cancel Edit
                    </button>
                    {/if}
                {/if}
                {#if !details.new}
                <button class="details-btn delete-btn" on:click={() => confirmDelete(details.data.id)}>Delete</button>
                {/if}
            </div>
            <div class="details-infos">
                <table>
                    {#if details.editable}
                        {#each Object.entries(details.data).filter(e => e[0] !== "id") as [k, v]}
                            <tr>
                                <td>{fieldToName(k)}</td>
                                <td id="{`details-${k}`}" contenteditable=true>
                                    {formatField(v)}
                                </td>
                            </tr>
                        {/each}
                    {:else}
                        {#each Object.entries(details.data).filter(e => e[0] !== "id" && e[1].length) as [k, v]}
                            <tr>
                                <td>{fieldToName(k)}</td>
                                <td>
                                    {#if k === "url"}
                                        <a href="#" on:click={() => Wails.Browser.OpenURL(v)}>{v}</a>
                                    {:else if Array.isArray(v)}
                                        {v.join(", ")}
                                    {:else}
                                        {v}
                                    {/if}
                                </td>
                            </tr>
                        {/each}
                    {/if}
                </table>
            </div>
        </div>
    {/if}
</div>