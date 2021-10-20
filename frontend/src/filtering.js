import {Document as FlexDocument} from "flexsearch";
import {statsStore} from "./stores";

let index;

export function buildFlexIndex(sequential_data) {
    index = new FlexDocument({
        tokenize: "full",
        document: {
            id: "id",
            index: [{
                field: "title",
            }, {
                field: "description",
            }, {
                field: "url",
            }, {
                field: "comment",
            }, {
                field: "resource_type",
                tokenize: "strict",
            }, {
                field: "tags",
                tokenize: "strict",
            }, {
                field: "regions",
                tokenize: "strict",
            }, {
                field: "licenses",
                tokenize: "strict",
            }, {
                field: "browsers",
                tokenize: "strict",
            }],
        }
    });
    for (let x = 0, data; x < sequential_data.length; x++) {
        data = sequential_data[x];
        index.add({
            id: data.id,
            title: data.title,
            description: data.description,
            url: data.url,
            comment: data.comment,
            resource_type: data.resource_type,
            licenses: data.licenses,
            browsers: data.browsers,
            tags: data.tags,
            regions: data.regions,
        }, {
            tokenize: "full",
        })
    }
}

export function filter(resources, filters) {
    // check special cases where index is not defined or filters is empty
    if (typeof index === 'undefined' || isEmpty(filters)) {
        return resources;
    }
    // prepare filters
    let searchfilters = []
    let limit = resources.length;
    for (let [fname, f] of filters.entries()) {
        if (f.length !== 0) {
            if (["resource_type", "tags", "regions", "licenses", "browsers"].includes(fname)) {
                f = f.join(" ");
            }
            searchfilters.push({field: fname, query: f, limit: limit});
        }
    }
    // execute flexsearch
    let results = index.search(searchfilters);
    // merge ids for matches
    // === AND ===
    if (results.length < searchfilters.length) {
        return [];
    }
    let matches = resources.map(r => r.id); // array of all resource ids as start value
    for (let result of results) {
        matches = matches.filter(value => result.result.includes(value)); // AND
    }
    // === OR ====
    // let matches = [];
    // for (let result of results) {
    //     matches = matches.concat(result.result); // OR
    // }
    // matches = matches.unique();
    // ===========
    // filter resources by match ids
    let filtered = resources.filter(elem => {
        if (matches.includes(elem["id"])) {
            return true;
        }
    });
    updateFilteredStats(searchfilters.length, filtered.length);
    return sortColumns(filtered);
}

function updateFilteredStats(filtersCount, filteredCount) {
    statsStore.update(s => {
        s.filteredCount = filteredCount;
        s.filtersCount = filtersCount;
        return s;
    })
}

Array.prototype.unique = function () {
    let a = this.concat();
    for (let i = 0; i < a.length; i++) {
        for (let j = i + 1; j < a.length; ++j) {
            if (a[i] === a[j]) a.splice(j--, 1)
        }
    }
    return a;
}

const isEmpty = m => {
    for (const e of m.values()) {
        if (e !== "") return false;
    }
    return true;
}

export function sortColumns(data, column=false) {
    if (column) {
        if (columnSort["column"] !== column) {
            columnSort["column"] = column;
            columnSort["sorted"] = false;
        }
        if (columnSort["sorted"]) {
            // already sorted, only needs to be reversed
            return data.reverse();
        }
    }
    column = columnSort["column"];
    const prepare = val => {
        // determine sort order per type
        // note: `undefined` will be sorted to the very end of the array
        switch (typeof val[column]) {
            case "string":
                // sort by lowercase strings in alphabetical order
                return val[column] !== "" ? val[column].toLowerCase() : undefined;
            case "object":
                // sort by number of elements
                return val[column].length !== 0 ? val[column].length : undefined;
            default:
                return val[column]
        }
    }
    data.sort((a, b) => {
        [a, b] = [prepare(a), prepare(b)];
        if (a === b) {
            return 0;
        } else if ((a < b && a !== undefined) || b === undefined) {
            return -1;
        } else {
            return 1;
        }
    });
    columnSort["sorted"] = true;
    return data;
}

let columnSort = {"column": "title", "sorted": false};

