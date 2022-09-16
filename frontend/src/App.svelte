<script>
  import { onMount } from "svelte"

  import SectionCard from "./components/SectionCard.svelte"

  import * as api from "./lib/api"

  /** @type {import("./lib/api").Devices} */
  let devices = []

  onMount(() => {
    console.log(`[App] [onMount] get devices...`)
    api.getDevices().then(res => devices = res)
  })
</script>

<main>
  <!-- Sections: -->
  <section
    style={`
      display: flex;
      flex-wrap: wrap;
      justify-content: center;
    `}
  >
    {#each devices as device}
      {#each device.sections as section}
        <div style="margin: 0.5rem;">
          <SectionCard
            host={device.host}
            port={device.port}
            sectionID={section.id}
          />
        </div>
      {/each}
    {/each}
  </section>
  <!-- Groups: -->
  <!-- Scenes: -->
</main>

<style lang="scss">
  @use "./sass/theme";

  main {
    overflow-x: hidden;

    & > section {
      margin: 1em;
    }
  }
</style>
