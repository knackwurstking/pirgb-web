<script>
  export let min = 0
  export let max
  export let value = 0
</script>

<label>
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
  }

  label span {
  }

  label input {
  }
</style>
