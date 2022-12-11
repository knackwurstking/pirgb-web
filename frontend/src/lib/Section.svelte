<script>
  /** @type {Section} */
  export let section;

  /** @type {boolean} */
  export let open = false;

  /** @type {boolean} */
  export let online = false;
</script>

<div class="section" class:open class:online />

<style>
  div.section {
    border: 1px solid red;
    width: 100%;
  }

  div.section.open {
    border: 1px solid green;
  }
</style>
