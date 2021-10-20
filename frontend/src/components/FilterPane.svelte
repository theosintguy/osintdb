<script>
    // TODO: hide scrollbar if height is enoough
    import {filtersStore, subscribeFilterPane as watch} from "../stores";
    import Tags from "svelte-tags-input";
    import {draw} from "svelte/transition";

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

    let resourceTypeAvailable;
    window.backend.DB.GetResourceTypes().then(r => {
        r.sort(); // TODO: necessary?
        resourceTypeAvailable = r;
    });

    let tagsAvailable; // TODO: reload when new tags have been added
    window.backend.DB.GetTags().then(t => {
        t.sort(); // TODO: necessary?
        tagsAvailable = t;
    })

    let regionsAvailable; // TODO: reload when new regions have been added
    window.backend.DB.GetRegions().then(r => {
        r.sort(); // TODO: necessary?
        regionsAvailable = r;
    })

    let browsersAvailable; // TODO: reload when new browsers have been added
    window.backend.DB.GetBrowsers().then(b => {
        b.sort(); // TODO: necessary?
        browsersAvailable = b;
    })

    let licensesAvailable; // TODO: reload when new licenses have been addded
    window.backend.DB.GetLicenses().then(l => {
        l.sort(); // TODO: necessary?
        licensesAvailable = l;
    })

    let show = false;
    watch(state => {
        show = state;
        if (show) {
            initLabels();
        }
    });

    $: filtersStore.set(filters);

    function initLabels() {
        setTimeout(() => {
            let labelName = "non-empty-input";
            let labels = document.querySelectorAll("div.field label");
            for (let label of labels) {
                let inputId = label.htmlFor;
                let inputElem = document.querySelector("input[id=" + inputId + "]");
                let isEmpty = (inputElem.value === "") && (inputElem.parentElement.childElementCount === 2);
                if (isEmpty) {
                    label.classList.remove(labelName);
                } else {
                    label.classList.add(labelName);
                }
            }
        }, 100);
    }

    function getFilterSetHandler(field) {
        return event => {
            filters.set(field, event.target.value);
            filters = filters; // trigger UI update
        }
    }

    function getFilterSetListHandler(field) {
        return event => {
            filters.set(field, event.detail.tags);
            filters = filters; // trigger UI update
        }
    }

    function setupFocusOutListeners() {
        let inputs = document.querySelectorAll(".container .field input");
        for (let elem of inputs) {
            elem.addEventListener("focusout", (event) => focusOutHandler(event.target));
        }
    }

    async function focusOutHandler(inputElem) {
        // check if current <input> is empty and no tags are active ?
        let isEmpty = (inputElem.value === "") && (inputElem.parentElement.childElementCount === 2);

        let inputId = inputElem.id;
        let label = document.querySelector(`label[for=${inputId}]`);

        if (isEmpty) {
            label.classList.remove("non-empty-input");
        } else {
            label.classList.add("non-empty-input");
        }
    }

</script>

<style>
    .container {
        height: 100%;
        width: 400px;
        box-shadow: 1px 0 5px -2px var(--bg-color-dark);
        display: flex;
        flex-direction: column;
        overflow-y: auto;
    }

    .title {
        margin: 0 0 10px;
        padding: 0;
        text-align: center;
    }

    form {
        display: flex;
        flex-direction: column;
    }

    .field,
    .field :global(.svelte-tags-input-layout) {
        position: relative;
    }

    .field input,
    .field :global(.svelte-tags-input) {
        width: 100%;
        padding: 4px 0;
        color: #fff;
        margin: 0 0 14px;
        border: none;
        border-bottom: 1px solid var(--bg-color2);
        outline: none;
        background: transparent;
        font-size: var(--font-size);
    }

    .field input:focus,
    .field :global(.svelte-tags-input:focus) {
        border-bottom-color: var(--primary-color);
    }

    .field label,
    .field :global(.svelte-tags-input-layout label) {
        position: absolute;
        top: 0;
        left: 0;
        padding: 4px 0;
        color: #fff;
        pointer-events: none;
        transition: .3s;
        font-size: var(--font-size);
    }

    input:focus ~ label,
    :global(form .field label.non-empty-input:not(:empty)), /* increased specificity */
    :global(.field .svelte-tags-input-layout:focus-within > label),
    :global(.field .svelte-tags-input-layout label.non-empty-input) {
        top: -19px;
        left: 0;
        color: var(--primary-color);
    }

    .field :global(.svelte-tags-input-layout),
    .field :global(.svelte-tags-input-layout):hover {
        background: transparent;
        outline: none;
        border: none;
        /*border-bottom: 1px solid var(--bg-color2);*/
        margin: 0;
        padding: 0;
    }

    .field :global(.svelte-tags-input-tag) {
        background: var(--bg-color-light-dark);
        margin: 0 0 10px;
        font-size: var(--font-size);
        border: none;
        border-bottom: 1px solid var(--bg-color2);
        border-radius: 0;
        padding: 0 0 4px;
    }

    .field :global(.svelte-tags-input-layout:focus-within > .svelte-tags-input-tag) {
        border-bottom: 1px solid var(--primary-color);
    }

    :global(.svelte-tags-input-matchs) {
        background: transparent;
    }

    .field :global(.svelte-tags-input-matchs) {
        background: var(--bg-color2);
        border: 2px solid var(--border-color);
        border-radius: 0;
        box-sizing: border-box;
        z-index: 1;
    }

    :global(::-webkit-scrollbar) {
        /*background: var(--bg-color3);*/
        width: 10px;
    }

    :global(::-webkit-scrollbar-thumb) {
        background: var(--bg-color);
    }

    :global(::-webkit-scrollbar-track) {
        background: var(--bg-color3);
    }

    .field :global(.svelte-tags-input-matchs li:hover),
    .field :global(.svelte-tags-input-matchs li:focus) {
        background: var(--bg-color);
    }

</style>

{#if show}
    <div class="container" use:setupFocusOutListeners>
        <h2 class="title">Filters</h2>
        <form autocomplete="off">
            <div class="field">
                <input id="titleInput" value={filters.get("title")}
                       on:input={getFilterSetHandler("title")}>
                <label for="titleInput">Title</label>
            </div>
            <div class="field">
                <Tags
                        id="resource_typeInput"
                        tags={filters.get("resource_type")}
                        on:tags={getFilterSetListHandler("resource_type")}
                        addKeys="{[188]}"
                        autoComplete="{resourceTypeAvailable}"
                        onlyAutocomplete
                        onlyUnique="{true}"
                        labelShow="{true}"
                        labelText="Type"
                        maxTags="1"
                />
            </div>
            <div class="field">
                <input id="descriptionInput" value={filters.get("description")}
                       on:input={getFilterSetHandler("description")}>
                <label for="descriptionInput">Description</label>
            </div>
            <div class="field">
                <input id="urlInput" value={filters.get("url")}
                       on:input={getFilterSetHandler("url")}>
                <label for="urlInput">URL</label>
            </div>
            <div class="field">
                <input id="commentInput" value={filters.get("comment")}
                       on:input={getFilterSetHandler("comment")}>
                <label for="commentInput">Comment</label>
            </div>
            <div class="field">
                <Tags
                        id="tagsInput"
                        tags={filters.get("tags")}
                        on:tags={getFilterSetListHandler("tags")}
                        addKeys="{[188]}"
                        autoComplete="{tagsAvailable}"
                        onlyAutocomplete
                        onlyUnique="{true}"
                        labelShow="{true}"
                        labelText="Tags"
                />
            </div>
            <div class="field">
                <Tags
                        id="regionsInput"
                        tags={filters.get("regions")}
                        on:tags={getFilterSetListHandler("regions")}
                        addKeys="{[188]}"
                        autoComplete="{regionsAvailable}"
                        onlyAutocomplete
                        onlyUnique="{true}"
                        labelShow="{true}"
                        labelText="Region"
                        maxTags="1"
                />
            </div>
            <div class="field">
                <Tags
                        id="licensesInput"
                        tags={filters.get("licenses")}
                        on:tags={getFilterSetListHandler("licenses")}
                        addKeys="{[188]}"
                        autoComplete="{licensesAvailable}"
                        onlyAutocomplete
                        onlyUnique="{true}"
                        labelShow="{true}"
                        labelText="License"
                        maxTags="1"
                />
            </div>
            <div class="field">
                <Tags
                        id="browsersInput"
                        tags={filters.get("browsers")}
                        on:tags={getFilterSetListHandler("browsers")}
                        addKeys="{[188]}"
                        autoComplete="{browsersAvailable}"
                        onlyAutocomplete
                        onlyUnique="{true}"
                        labelShow="{true}"
                        labelText="Browser"
                        maxTags="1"
                />
            </div>
        </form>
    </div>
{/if}