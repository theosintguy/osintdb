import {writable} from 'svelte/store';
import Wails from '@wailsapp/runtime';
import debounce from 'lodash.debounce';

// ----- Simple Stores -----

// filterPaneDisplay
const showFilterPane = writable(false);
export const subscribeFilterPane = showFilterPane.subscribe;

export function toggleFilterPane() {
    showFilterPane.update(state => !state);
}

// add resource button
const addResources = writable(0); // hack
export const subscribeAddResources = addResources.subscribe;

export function addResource() {
    addResources.update(count => count + 1);
}

// stats store
export const statsStore = writable({
    "resourcesCount": 0,
    "filteredCount": 0,
    "filtersCount": 0,
});

// ----- Sync Stores -----

// tableDataStore
export let resourcesStore;

export function initStores() {
    resourcesStore = Wails.Store.New("resourcesStore");
}


// ----- Custom Stores -----

/**
 * DelayStore works similar to Writable Stores. The only difference is the `set` method.
 * The `set` method only updates the store's value after a delay. If in the meantime other
 * calls to `set` have occurred it will update the store's value to the latest of these values.
 * Other `set` calls that happen during the delay time will be blocked using the mutex (and then cancelled).
 */
class DelayStore {
    constructor(initialValue, delay = 0, maxDelay = 0) {
        this.newSetValue = initialValue;
        this.newUpdateValue = initialValue;
        this.store = writable(initialValue);
        this.subscribe = this.store.subscribe;
        // this.update = this.store.update;
        const debounceOptions = maxDelay !== 0 ? {maxWait: maxDelay} : {};
        this.commit = debounce(() => {
            this.store.set(this.newSetValue);
        }, delay, debounceOptions);
    }

    async set(value) {
        this.newSetValue = value;
        this.commit();
    }
}

// filters
export const filtersStore = new DelayStore(new Map(), 400, 1500);
