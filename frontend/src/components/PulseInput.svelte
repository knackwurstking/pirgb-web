<script>
  export let min = undefined
  export let max = undefined
  export let value
</script>

<label {...$$props}>
  <span>Pulse</span>
  <input
    type="number"
    {min}
    {max}
    bind:value
  />
</label>

<style lang="scss">
  @use "../sass/theme";

  $border-width: theme.$border-width;
  $border-style: theme.$border-style;
  $border-color: theme.$border-color;

  label {
    display: flex;
    flex-direction: column;
    place-items: center;
    width: 9.5rem;
    height: 2.5rem;
    max-width: 10rem;
    padding: 0.1em;
  }

  label span {
    font-size: 0.70rem;
    width: 100%;
    text-align: center;
  }

  label input {
    text-align: center;
    width: 5rem;
    border-top: $border-width $border-style $border-color;
    border-right: none;
    border-left: none;
  }
</style>
