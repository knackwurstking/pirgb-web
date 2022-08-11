<script>
  import {
    Popover,
    PopoverButton,
    PopoverPanel,
  } from "@rgossiaux/svelte-headlessui"

  import {
    ChevronDownIcon,
  } from "@rgossiaux/svelte-heroicons/solid"

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
      href="/schemes/light.css"
    />
  {:else if scheme === "dark"}
    <link
      rel="stylesheet"
      href="/schemes/dark.css"
    />
  {:else}
    <link
      rel="stylesheet"
      href="/schemes/light.css"
      media="(prefers-color-scheme: light)"
    />

    <link
      rel="stylesheet"
      href="/schemes/dark.css"
      media="(prefers-color-scheme: dark)"
    />
  {/if}
</svelte:head>

<Popover
  style={`
     position: absolute;
     top: 0;
     left: 0;
   `}
  let:close
>
  <PopoverButton
    style={`
      display: flex;
      align-items: center;
      justify-content: space-between;
      width: 10rem;
      position: absolute;
      top: 1rem;
      left: 2rem;
    `}
  >
    <span>{selectedItem.name}</span>
    <ChevronDownIcon style="width:1.5rem;"/>
  </PopoverButton>

  <PopoverPanel
    style={`
      position: absolute;
      z-index: 999;
      background-color: var(--color-bg-secondary);
      border-radius: var(--radius);
      padding: 1rem;
      border: 0.1rem solid;
      border-color: var(--color-border);
      width: 25rem;
      top: 4.5rem;
      left: 1.5rem;
    `}
  >
    <div class="popover-panel-content" transition:slide>
      {#each items as item, index (item.id)}
        <a class="popover-panel-item" href={item.href} on:click={() => {
           selectedItem = items[index];
           close(null);
        }}>{item.name}</a>
      {/each}
    </div>
  </PopoverPanel>
</Popover>

<style>
  .popover-panel-content {
    display: flex;
    flex-direction: column;
  }

  .popover-panel-content > * {
    margin: 0.5rem 0;
  }

  .popover-panel-item {
    border-radius: var(--radius);
    /*
    border: 0.1rem solid;
    border-color: var(--color-border);
    */
    padding: 0.5rem 0;
  }

  .popover-panel-item:hover,
  .popover-panel-item:active {
    /*
    border-color: var(--color-base);
    */
    background-color: var(--color-bg-highlight);
  }
</style>
