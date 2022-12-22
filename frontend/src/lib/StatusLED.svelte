<script>
    export let active = false;
    export let activeColor = "green";
    export let inactiveColor = "red";
</script>

<div class="status-led--outer" class:active {...$$restProps}>
    <div
        style={`
            background-color: ${active ? activeColor : inactiveColor};
        `}
        class="status-led--inner"
    />
</div>

<style>
    .status-led--outer {
        position: relative;
        padding: 4px;
        border: 1px solid var(--color-border);
        border-radius: 50%;
        width: calc(1rem + 8px);
        height: calc(1rem + 8px);
    }

    .status-led--outer.active .status-led--inner {
        border-radius: 50%;
        width: 100%;
        height: 100%;
        filter: blur(0.125rem);
        /* TODO: pulse animation blur value up and down */
    }
</style>
