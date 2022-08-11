<script>
    import {
        Disclosure,
        DisclosurePanel,
        DisclosureButton,
    } from "@rgossiaux/svelte-headlessui";

    import {
        ChevronRightIcon,
    } from "@rgossiaux/svelte-heroicons/solid";

    import { fade } from "svelte/transition";

    let width = "40rem";

    let sections = [
        { id: 0, title: "Sections", open: true },
        { id: 1, title: "Groups", open: false },
        { id: 2, title: "Scripts", open: false },
    ];
</script>

{#each sections as section (section.id)}
  <Disclosure style={`width:${width};margin:0.5rem 0;`} let:open defaultOpen={section.open}>
    <DisclosureButton style="width:100%;margin:0.5rem 0;">
      <div style="display:flex;position:relative;align-items:center;justify-content:space-between;">
        <div style="padding-right: 1rem">{section.title}</div>
        <ChevronRightIcon
            style={(open ? "transform: rotate(90deg);": "") + "width: 1.5rem"}
        />
      </div>
    </DisclosureButton>
    {#if open}
      <div transition:fade>
        <DisclosurePanel style="width:100%;margin: 0.5rem 0;padding:0.5rem">
          <span>... here should be a list with {section.title.toLowerCase()} ...</span>
        </DisclosurePanel>
      </div>
    {/if}
  </Disclosure>
{/each}

<style>
</style>
