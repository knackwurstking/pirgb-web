<script>
  export let min
  export let max
  export let value = 0
</script>

<label {...$$props}>
  <span>{value}</span>
  <input
    type="range"
    {min}
    {max}
    bind:value
  />
</label>

<style>
  label {
    border: 1px solid red;
    display: flex;
    flex-wrap: nowrap;
    flex-direction: row-reverse;
  }

  label span {
    margin: 0.25em;
  }

  label input {
    margin: 0.25em;
  }
</style>
