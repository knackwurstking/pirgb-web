<script>
  import { Popover, PopoverButton, PopoverPanel } from "@rgossiaux/svelte-headlessui"
  import { ColorSwatchIcon } from "@rgossiaux/svelte-heroicons/solid"
  import { onMount } from "svelte"
  import { slide } from "svelte/transition"

  export let schemes = [
    "System",
    "Light",
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
  {#if currentScheme === "Light" }
    <link
      rel="stylesheet"
      href="/schemes/light.css"
    />
  {:else if currentScheme === "Dark"}
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
      top: 0.5em;
      right: 0.5em;
    `}
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
      padding: 0 0.25em;
      width: fit-content;
      top: 3.5em;
      right: 0.5em;
    `}
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
    width: 7em;
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
