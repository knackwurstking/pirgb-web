<script>
  import { onMount } from "svelte";

  onMount(() => {
    console.log("/settings");
  });
</script>

<div class="settings">
</div>

<style>
  .settings {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
</style>
