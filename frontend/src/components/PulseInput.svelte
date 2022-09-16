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

  $borderWidth: 0.1rem;
  $borderStyle: solid;
  $borderColor: theme.$divider;

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
    border-top: $borderWidth $borderStyle $borderColor;
    border-right: none;
    border-left: none;
  }
</style>
