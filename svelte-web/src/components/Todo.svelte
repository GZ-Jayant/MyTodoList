<script>
    import TodoItem from "./TodoItem.svelte";
    import { onMount } from "svelte";

    let todoList = [];
    export let currentFilter = "all";
    let filteredTasks = [];
    let slicedTasks = [];
    let isAddingTask = false; // 控制按钮显示和隐藏的状态变量
    let currentPage = 1;
    const itemsPerPage = 3;
    export let apiUrl;
    console.log("api: " + apiUrl);

    /**
     * 获取当前日期，n为偏移天数
     * @param {number} [n=0] - 偏移天数
     * @returns {string} - 格式化后的日期字符串
     */
    function getCurrentDate(n = 0) {
        let date = new Date();
        date.setDate(date.getDate() + n);
        const year = date.getFullYear();
        const month = String(date.getMonth() + 1).padStart(2, "0"); // 月份从0开始，需要加1
        const day = String(date.getDate()).padStart(2, "0");
        return `${year}-${month}-${day}`;
    }

    /**
     * 添加新任务
     */
    function addTask() {
        const newTask = {
            task_id: -1,
            task_name: "",
            task_desc: "",
            task_status: "active",
            task_start_date: getCurrentDate(),
            task_end_date: getCurrentDate(1),
        };
        currentFilter = "all";
        todoList = [...todoList, newTask];
        isAddingTask = true;
    }

    /**
     * 查询任务列表
     * @param {string} [filter="all"] - 过滤条件
     */
    function selectTask(filter = "all") {
        let formData = new URLSearchParams();
        formData.append("task_status", filter);

        const fetchPromise = fetch(apiUrl + "/task/select", {
            method: "POST",
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            },
            body: formData,
        });

        fetchPromise
            .then((response) => {
                if (!response.ok) {
                    throw new Error(`HTTP 请求错误：${response.status}`);
                }
                return response.json();
            })
            .then((data) => {
                todoList = data["taskList"];
            })
            .catch((error) => {
                console.log(`无法获取产品列表：${error}`);
            });
    }

    $: filteredTasks = todoList
        .filter((task) => {
            if (currentFilter === "all") {
                return true;
            } else return task.task_status === currentFilter;
        })
        .sort((a, b) => {
            // 特殊条件：ID为-1的任务置顶
            if (a.task_id === -1) return -1;
            if (b.task_id === -1) return 1;
            // 其他任务按日期倒序排列
            const dateComparison =
                new Date(b.task_start_date) - new Date(a.task_start_date);
            if (dateComparison !== 0) {
                return dateComparison;
            }

            // 如果日期相等，按ID倒序排列
            return b.task_id - a.task_id;
        });
    $: slicedTasks = filteredTasks.slice(
        (currentPage - 1) * itemsPerPage,
        currentPage * itemsPerPage,
    );
    $: canGoToNextPage = currentPage * itemsPerPage < filteredTasks.length;

    // 处理任务编辑完成或退出编辑的事件
    function handleTaskUpdate(event) {
        if (event.detail === "create") {
            isAddingTask = false; // 显示按钮
        }
        selectTask(currentFilter);
    }

    // 翻页函数
    function nextPage() {
        if (canGoToNextPage) {
            currentPage += 1;
        }
    }

    function prevPage() {
        if (currentPage > 1) {
            currentPage -= 1;
        }
    }

    onMount(() => {
        selectTask();
    });
</script>

{#if !isAddingTask}
    <button class="add-task" on:click={addTask}>Add New Task</button>
{/if}
<ul class="list-container">
    {#each slicedTasks as todo}
        <TodoItem {todo} {apiUrl} on:changeTodo={handleTaskUpdate} />
    {:else}
        <li>Nothing to do here!</li>
    {/each}
</ul>
<div class="pagination">
    <button on:click={prevPage} disabled={currentPage === 1}>上一页</button>
    <span>第 {currentPage} 页</span>
    <button on:click={nextPage} disabled={!canGoToNextPage}>下一页</button>
</div>

<style>
    .add-task {
        width: 150px; /* 固定宽度 */
        height: 45px; /* 固定高度 */
        margin-bottom: 0.5rem; /* 增加底部间距 */
    }
    .list-container {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        font-family: "Consolas", monospace;
    }
    .pagination {
        display: flex;
        justify-content: space-between;
        margin-top: 10px;
    }
</style>
