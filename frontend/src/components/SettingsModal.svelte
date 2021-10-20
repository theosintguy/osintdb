<script>
    import {getContext} from 'svelte';
    const {close} = getContext('simple-modal');

    function _import() {
        window.backend.DB.ImportData(); // TODO: display possible error message to user
        
    }
    function _export() {
        window.backend.DB.ExportData(); // TODO: display possible error message to user
    }
    function _clearDB() {
        window.backend.DB.ClearData(); // TODO: display possible error message to user
    }
</script>

<style>
    :global(div.window-wrap div.window) {
        background: var(--bg-color2);
        color: var(--text-color);
        border-radius: var(--border-radius);
    }

    :global(div.window-wrap div.window button.close) {
        display: none;
    }

    div.btns {
        display: flex;
        flex-direction: row;
        justify-content: center;
        gap: 16px;
        margin-top: 16px;
    }

    button {
        border: var(--border);
        border-radius: var(--border-radius);
        box-shadow: 1px 1px 2px var(--bg-color-dark);
        padding: 6px;
        cursor: pointer;
        background: var(--bg-color2);
        color: var(--text-color);
    }

    button:hover {
        background: var(--bg-color3);
    }

    button:active {
        background: var(--border-color);
        box-shadow: none;
    }
</style>

<button on:click={close}>Go back</button>
<div>
    <h2>Import/Export</h2>
    Here you can import or export tools to the Database using JSON files. Please note that imported data will not be
    checked against the Database and could thereby lead to duplicates. This will be resolved in a future release.
    Until then it is recommended to first export the current data and use that to dedupe the import data beforehand.
    <div class="btns">
        <button on:click={_import}>Import</button>
        <button on:click={_export}>Export</button>
    </div>
</div>
<div>
    <h2>Clear Data</h2>
    <b>Attention: This will delete all your data. Make sure to back it up before.</b>
    <br />
    Also note that osintDB will try to load the default data again at the next startup if it finds an empty database.
    <div class="btns">
        <button on:click={_clearDB}>Delete</button>
    </div>
</div>
