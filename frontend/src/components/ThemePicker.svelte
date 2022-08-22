<script>
  import { Popover, PopoverButton, PopoverPanel } from "@rgossiaux/svelte-headlessui"
  import { ColorSwatchIcon } from "@rgossiaux/svelte-heroicons/solid"
  import { onMount } from "svelte"
  import { slide } from "svelte/transition"

  export let schemes = [
    "system",
    "mono dark",
    "mono light",
    "dark blue",
  ]

  // TODO: Store state in (window) local storage
  let mounted = false
  export let currentScheme = schemes[0];
  $: (mounted) && window.localStorage.setItem("scheme", currentScheme)

  onMount(() => {
    currentScheme = window.localStorage.getItem("scheme") || "system"
    mounted = true
  })
</script>

<svelte:head>
  {#if currentScheme === "mono light" }
    <link
      rel="stylesheet"
      href="/schemes/mono-light.css"
    />
  {:else if currentScheme === "mono dark"}
    <link
      rel="stylesheet"
      href="/schemes/mono-dark.css"
    />
  {:else if currentScheme === "dark blue"}
    <link
      rel="stylesheet"
      href="/schemes/dark-blue.css"
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
      background-color: var(--color-bg-secondary);
      border: 0.1rem solid var(--border-color);
      padding: 0 0.25em;
      width: 10em;
      top: 3.5em;
      right: 0.5em;
    `}
  >
    <div
      style={`
        display: flex;
        flex-direction: column;
        place-items: center;
      `}
      transition:slide
    >
      <!-- TODO: list schemes available here ... -->
      {#each schemes as scheme, index}
        <div
          class="panel-item"
          on:click={() => {
            currentScheme = schemes[index]
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
  .panel-item {
    margin: 0.25em 0;
    font-weight: 500;
    color: var(--color);
    text-decoration: inherit;
    padding: 0.25em 0;
    transition: background-color .5s ease, color .5s ease;
    width: 100%;
  }

  .panel-item:hover,
  .panel-item:active {
    color: var(--color-primary-bright);
    background-color: var(--color-bg-highlight);
  }
</style>
