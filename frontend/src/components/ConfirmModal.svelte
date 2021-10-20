<script>
    import {getContext} from 'svelte';
    const {close} = getContext('simple-modal');

    export let message;
    export let onCancel = () => {};
    export let onOkay = () => {};

    function _onCancel() {
       onCancel();
       close();
    }

    function _onOkay() {
        onOkay();
        close();
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

<p>
    {message}
</p>
<div class="btns">
    <button on:click={_onOkay}>Ok</button>
    <button on:click={_onCancel}>Cancel</button>
</div>