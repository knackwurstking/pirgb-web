<script>
  import { onMount } from "svelte";

  export let navigate;

  onMount(() => {
    if (navigate) navigate("/sections");
  });
</script>
