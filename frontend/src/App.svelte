<script>
    import {
        Popover,
        PopoverButton,
        PopoverPanel,
    } from "@rgossiaux/svelte-headlessui";

    import {
        ChevronDownIcon,
    } from "@rgossiaux/svelte-heroicons/solid";

    import { slide } from "svelte/transition";

    //let width = "40rem";

    let items = [
        { id: 0, name: "Sections" },
        { id: 1, name: "Groups" },
        { id: 2, name: "Scripts" },
    ];
    let selectedItem = items[0]
</script>

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
      background-color: var(--secondary-bg-color);
      border-radius: var(--border-radius);
      padding: 1rem;
      border: 0.1rem solid;
      border-color: var(--secondary-color);
      width: 25rem;
      top: 4.5rem;
      left: 1.5rem;
    `}
  >
    <!-- TODO: do some slide in transition from top to bottom -->
    <div class="popover-panel-content" transition:slide>
      {#each items as item (item.id)}
        <a href="#" on:click={() => close()}>{item.name}</a>
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
</style>
