<script>
  import {
    Popover,
    PopoverButton,
    PopoverPanel,
  } from "@rgossiaux/svelte-headlessui"

  //import {
  //  ChevronDownIcon,
  //} from "@rgossiaux/svelte-heroicons/solid"

  import { onMount } from "svelte"
  import { slide } from "svelte/transition"

  //import ThemePicker from "./components/ThemePicker.svelte"
  import SectionCard from "./components/SectionCard.svelte"
  import FlyDiv from "./components/FlyDiv.svelte"

  import * as api from "./lib/api"

  let items = [
    "Sections",
    "Groups",
    "Scenes",
  ]

  let selectedItem = items[0]

  /** @type {import("./lib/api").Devices} */
  let devices = []

  onMount(() => {
    console.log(`[App.svelte] [onMount] App.svelte`)
    api.getDevices().then(res => devices = res)
  })
</script>

<svelte:head>
  <link
    rel="stylesheet"
    href="/schemes/dark.css"
  />
</svelte:head>

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
      width: 10em;
      position: absolute;
      top: 0.5em;
      left: 0.5em;
    `}
  >
    <span>{selectedItem}</span>
  </PopoverButton>

  <PopoverPanel
    style={`
      position: absolute;
      z-index: 10;
      background-color: var(--button-bg);
      border: var(--border-width) var(--border-style) var(--border-color);
      padding: 0 0.25em;
      width: fit-content;
      top: 3.5em;
      left: 0.5em;
    `}
  >
    <div 
      class="panel-item-container"
      transition:slide
    >
      {#each items as item}
        <div
          class="panel-item"
          on:click={() => {
           selectedItem = item;
           close(null);
          }}
        >
          {item}
        </div>
      {/each}
    </div>
  </PopoverPanel>
</Popover>

<main>
  {#if selectedItem === "Sections"}
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
          <div style="margin: 0.5em;">
            <SectionCard
              host={device.host}
              port={device.port}
              sectionID={section.id}
            />
          </div>
        {/each}
      {/each}
    </FlyDiv>
  {/if}
</main>

<style>
  .panel-item-container {
    display: flex;
    flex-direction: column;
    width: 10em;
  }

  .panel-item {
    margin: 0.25em 0;
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
    background-color: var(--alternate-bg-color);
  }
</style>
