<script>
    export let active = false;
    export let activeColor = "green";
    export let inactiveColor = "red";
</script>

<div class="status-led-container" class:active>

</div>

<style>
    /* TODO: inactive styles */
    div.status-led-container {
        position: relative;
        border: 50%;
    }

    /* TODO: active styles */
    div.status-led-container.active {
    }
</style>
