<script>
  import { onMount } from "svelte";

  onMount(() => {
    console.log("/settings");
  });
</script>

<div class="settings">
  <section class="authorization">
    <div class="api-key">
      <!-- TODO: classic label and input for api key (and a status led for validation?) -->
      <label for="api-key">Api Key:</label><br />
      <input type="text" name="api-key" />
    </div>
  </section>
</div>

<style>
  .settings {
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .settings .authorization {
  }

  .settings .authorization .api-key {
    padding: 8px;
    margin: 8px;
  }

  .settings .authorization .api-key label {
    margin: 8px;
  }

  .settings .authorization .api-key input {
    margin-top: 8px;
  }
</style>
