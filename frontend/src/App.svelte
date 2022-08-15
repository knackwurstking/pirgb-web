<script lang="js">
  import {
    Popover,
    PopoverButton,
    PopoverPanel,
  } from "@rgossiaux/svelte-headlessui"

  import Paper, {
      Title as PaperTitle,
      Content as PaperContent,
  } from "@smui/paper"

    import { getSections } from "./lib/api"

  //import {
  //  ChevronDownIcon,
  //} from "@rgossiaux/svelte-heroicons/solid"

  import { slide } from "svelte/transition"

  export let scheme = "";

  let items = [
    { id: 0, href: "#", name: "Sections" },
    { id: 1, href: "#", name: "Groups" },
    { id: 2, href: "#", name: "Scene" },
  ]

  let selectedItem = items[0]
  $: {
    if (selectedItem)
        console.info(`Selected table: "${selectedItem.name}"`)
  }
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

<Popover
  style={`
    position: absolute;
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
      width: 9rem;
      position: absolute;
      top: 0.5rem;
      left: 0.5rem;
    `}
  >
    <span>{selectedItem.name}</span>
    <!--
    <ChevronDownIcon style="width:1.5rem;"/>
    -->
  </PopoverButton>

  <PopoverPanel
    style={`
      position: absolute;
      z-index: 10;
      background-color: var(--bg-mid);
      border-radius: var(--radius);
      padding: 0 0.25rem;
      border: 0.1rem solid;
      border-color: var(--border-color, --bg-low);
      width: 15rem;
      top: 3.5rem;
      left: 0.5rem;
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
  <!-- 
    TODO: render (sections/groups/scene) content here
    - transition if item changed: old swipe out (right to left) and new swipe in (right to left)
    - column layout centered, top to bottom
  -->
  {#each getSections() as section (section.id)}
    <!-- TODO: Create a Paper for each section (host, port, sectionID, quick on/off buttons or just a toggle switch) -->
  {/each}
</main>

<style lang="css">
  .popover-panel-content {
    display: flex;
    flex-direction: column;
  }

  .popover-panel-content > * {
    margin: 0.25rem 0;
  }

  .popover-panel-item {
    border-radius: var(--radius);
    transition: background-color .5s ease, color .5s ease;
    padding: 0.5rem 0;
  }

  .popover-panel-item:hover,
  .popover-panel-item:active {
    background-color: var(--bg-low);
  }
</style>
