<script>
	import Todo from "./components/Todo.svelte";
	import Filter from "./components/Filter.svelte";

	let currentFilter = "all";
</script>

<main>
	<h1>Todo List</h1>
	<div class="main-ui">
		<div class="filter-wrapper">
			<Filter bind:currentFilter />
		</div>
		<div class="todo-wrapper">
			<Todo {currentFilter} />
		</div>
	</div>
</main>

<style>
	:global(body) {
		background-color: #959595;
	}
	.main-ui {
		display: flex;
		justify-content: center;
		align-items: flex-start;
	}
	.filter-wrapper {
		padding: 20px;
	}
	.todo-wrapper {
		padding: 20px;
		margin-right: 20px; /* 保持与filter-wrapper有间距 */
	}
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #72227f;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>
