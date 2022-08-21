<script>
  import { fly } from "svelte/transition"
</script>

<div
  {...$$props}
  in:fly="{{ y: -1000, duration: 500, delay: 250 }}"
  out:fly="{{ y: 1000, duration: 500 }}"
>
  <slot />
</div>
