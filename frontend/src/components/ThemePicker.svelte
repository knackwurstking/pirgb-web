<script>
  import { Popover, PopoverButton, PopoverPanel } from "@rgossiaux/svelte-headlessui"
  import { ColorSwatchIcon } from "@rgossiaux/svelte-heroicons/solid"
  import { onMount } from "svelte"
  import { slide } from "svelte/transition"

  export let schemes = [
    "Dark Blue",
    "Dark White",
    "Dark Green",
    "Green",
    "Dark",
  ]

  let mounted = false
  export let currentScheme = schemes[0];
  $: (mounted) && window.localStorage.setItem("scheme", currentScheme)

  onMount(() => {
    currentScheme = window.localStorage.getItem("scheme") || "system"
    mounted = true
  })
</script>

<svelte:head>
  {#if currentScheme === "Dark Blue" }
    <link
      rel="stylesheet"
      href="/schemes/dark-blue.css"
    />
  {:else if currentScheme === "Dark White"}
    <link
      rel="stylesheet"
      href="/schemes/dark-white.css"
    />
  {:else if currentScheme === "Dark Green"}
    <link
      rel="stylesheet"
      href="/schemes/dark-green.css"
    />
  {:else if currentScheme === "Green"}
    <link
      rel="stylesheet"
      href="/schemes/green.css"
    />
  {:else if currentScheme === "Dark"}
    <link
      rel="stylesheet"
      href="/schemes/dark.css"
    />
  {:else}
    <link
      rel="stylesheet"
      href="/schemes/dark-blue.css"
    />
  {/if}
</svelte:head>

<Popover
  style={`
    position: fixed;
    z-index: 998;
    top: 0;
    right: 0;
  `}
  let:close
>
  <PopoverButton
    style={`
      position: absolute;
      top: 0.5rem;
      right: 1rem;
    `}
    class="elevate-900"
  >
    <ColorSwatchIcon
      style={`
        width: 1rem;
        height: 1rem;
      `}
    />
  </PopoverButton>

  <PopoverPanel
    style={`
      position: absolute;
      z-index: 10;
      background-color: var(--button-bg);
      border: var(--border-width) var(--border-style) var(--border-color);
      padding: 0 0.25rem;
      width: fit-content;
      top: 3.5rem;
      right: 1rem;
    `}
    class="elevate-900"
  >
    <div
      class="panel-item-container"
      transition:slide
    >
      {#each schemes as scheme}
        <div
          class="panel-item"
          on:click={() => {
            currentScheme = scheme
            close(null)
          }}
        >
          {scheme}
        </div>
      {/each}
    </div>
  </PopoverPanel>
</Popover>

<style>
  .panel-item-container {
    display: flex;
    flex-direction: column;
    width: 7rem;
  }

  .panel-item {
    margin: 0.25rem 0;
    font-size: inherit;
    font-weight: 500;
    text-decoration: inherit;
    padding: 0.35rem 0.6rem;
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
