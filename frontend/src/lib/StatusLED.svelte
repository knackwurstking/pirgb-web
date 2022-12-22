<script>
    export let active = false;
    export let activeColor = "green";
    export let inactiveColor = "red";
</script>

<div class="status-led--outer" class:active {...$$restProps}>
    <span
        style={`
            background-color: ${active ? activeColor : inactiveColor};
        `}
        class="status-led--inner"
    />
</div>

<style>
    /* inactive styles */
    .status-led--outer {
        position: relative;
        border: 50%;
        padding: 8px;
        border: 1px solid var(--color-border);
        width: 1rem;
        height: 1rem;
    }

    /* active styles */
    .status-led--outer.active .status-led--inner {
        width: 100%;
        height: 100%;
    }
</style>
