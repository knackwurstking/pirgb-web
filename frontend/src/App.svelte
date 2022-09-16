<script>
  import {
    Popover,
    PopoverButton,
    PopoverPanel,
  } from "@rgossiaux/svelte-headlessui"

  import { onMount } from "svelte"

  import SectionCard from "./components/SectionCard.svelte"
  import FlyDiv from "./components/FlyDiv.svelte"

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
  <FlyDiv
    style={`
      display: flex;
      flex-wrap: wrap;
      justify-content: center;
      padding-top: 3em;
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
  </FlyDiv>
  <!-- Groups: -->
  <!-- Scenes: -->
</main>

<style lang="scss">
  @use "./sass/theme";

  main {
    overflow-x: hidden;
  }

  .panel-item-container {
    display: flex;
    flex-direction: column;
    width: 10rem;
  }

  .panel-item {
    margin: 0.25rem 0;
    font-size: inherit;
    font-weight: 500;
    text-decoration: inherit;
    padding: 0.35em 0.6em;
    transition: background-color .5s ease, color .5s ease;
    width: 100%;
    text-align: center;
    cursor: pointer;
  }

  .panel-item:hover,
  .panel-item:active {
    background-color: theme.$surface;
  }
</style>
