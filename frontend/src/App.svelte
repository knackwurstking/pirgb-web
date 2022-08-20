<script>
  import {
    Popover,
    PopoverButton,
    PopoverPanel,
    Transition
  } from "@rgossiaux/svelte-headlessui"

  //import {
  //  ChevronDownIcon,
  //} from "@rgossiaux/svelte-heroicons/solid"

  import { onMount } from "svelte"
  import { slide } from "svelte/transition"

  import SectionCard from "./components/SectionCard.svelte"

  import * as api from "./lib/api"

  export let scheme = "";

  /** @type {{ id: number, href: string, name: string }[]} */
  let items = [
    { id: 0, href: "#sections", name: "Sections" },
    { id: 1, href: "#groups", name: "Groups" },
    { id: 2, href: "#scenes", name: "Scenes" },
  ]

  // TODO: on mount - parse hash (window.location) an set selected Item
  /** @type {{ id: number, href: string, name: string }} */
  let selectedItem = items[0]

  /** @type {import("./lib/api").Devices} */
  let devices = []

  onMount(() => {
    console.log(`[onMount] App.svelte`)
    api.getDevices().then(res => devices = res)
  })
</script>

<svelte:head>
  {#if scheme === "light" }
    <link
      rel="stylesheet"
      href="/schemes/mono-light.css"
    />
  {:else if scheme === "dark"}
    <link
      rel="stylesheet"
      href="/schemes/mono-dark.css"
    />
  {:else}
    <link
      rel="stylesheet"
      href="/schemes/mono-light.css"
      media="(prefers-color-scheme: light)"
    />

    <link
      rel="stylesheet"
      href="/schemes/mono-dark.css"
      media="(prefers-color-scheme: dark)"
    />
  {/if}
</svelte:head>

<!-- TODO: Add a theme picker on the top right position -->
<Popover
  style={`
    position: fixed;
    z-index: 999;
    top: 0;
    left: 0;
   `}
  let:close
>
  <PopoverButton
    style={`
      display: flex;
      align-items: center;
      justify-content: center;
      width: 9em;
      position: absolute;
      top: 0.5em;
      left: 0.5em;
    `}
  >
    <span>{selectedItem.name}</span>
    <!--
    <ChevronDownIcon style="width:1.5em;"/>
    -->
  </PopoverButton>

  <PopoverPanel
    style={`
      position: absolute;
      z-index: 10;
      background-color: var(--color-bg-secondary);
      border-radius: var(--border-radius);
      padding: 0 0.25em;
      border: 0.1em solid;
      border-color: var(--border-color);
      width: 15em;
      top: 3.5em;
      left: 0.5em;
    `}
  >
    <div class="popover-panel-content" transition:slide>
      {#each items as item, index (item.id)}
        <a
          class="popover-panel-item"
          href={item.href},
          on:click={() => {
           selectedItem = items[index];
           close(null);
          }}
        >
          {item.name}
        </a>
      {/each}
    </div>
  </PopoverPanel>
</Popover>

<main>
  <!-- TODO: Transition out: up direction, Transition in: from the button -->
  <Transition
    show={selectedItem.name.toLowerCase() === "sections"}
  >
    <div class="flex-container">
      {#each devices as device}
        {#each device.sections as section}
          <div class="item">
            <SectionCard
              host={device.host}
              port={device.port}
              sectionID={section.id}
            />
          </div>
        {/each}
      {/each}
    </div>
  </Transition>
</main>

<style>
  .popover-panel-content {
    display: flex;
    flex-direction: column;
  }

  .popover-panel-content > * {
    margin: 0.25em 0;
  }

  .popover-panel-item {
    border-radius: var(--border-radius);
    transition: background-color .5s ease, color .5s ease;
    padding: 0.5em 0;
  }

  .popover-panel-item:hover,
  .popover-panel-item:active {
    background-color: var(--color-bg-highlight);
  }

  div.flex-container {
      position: absolute;
      top: 0;
      right: 0;
      bottom: 0;
      left: 0;
      display: flex;
      flex-wrap: wrap;
      justify-content: center;
      padding-top: 3rem;
  }

  div.flex-container .item {
    margin: 0.5rem;
  }
</style>
