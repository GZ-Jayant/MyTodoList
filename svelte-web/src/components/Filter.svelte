<script>
    export let currentFilter = "all";
</script>

<form>
    <select bind:value={currentFilter}>
        <option value="all">All</option>
        <option value="active">Active</option>
        <option value="overdue">Overdue</option>
        <option value="completed">Completed</option>
    </select>
</form>

<style>
    select {
        margin: 10px;
    }
</style>
