<script>
  import { onMount } from "svelte"
  import SectionCard from "./lib/SectionCard.svelte"
  import api from "./js/api"

  /** @type {import("./js/api").Devices} */
  let devices = []

  onMount(() => {
    console.log(`[App.svelte] onMount: get devices...`)
    api.getDevices().then(res => devices = res)
  })
</script>

<main>
  <section class="sections">
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
</main>

<style>
  main {
    overflow-x: hidden;
  }

  section {
    margin: 1rem;
  }
</style>
