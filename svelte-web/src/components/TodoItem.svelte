<script>
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();
    export let apiUrl;

    export let todo = {
        task_id: 0,
        task_name: "",
        task_desc: "",
        task_status: "",
        task_start_date: "",
        task_end_date: "",
    };

    let backupTodo = { ...todo };
    let isEditing = false;
    $: {
        if (todo.task_id === -1) {
            isEditing = true;
        }
    }

    function createTask({
        tname = todo.task_name,
        tdesc = todo.task_desc,
        tstatus = todo.task_status,
        tstart = todo.task_start_date,
        tend = todo.task_end_date,
    }) {
        let formData = new URLSearchParams();
        formData.append("task_name", tname);
        formData.append("task_desc", tdesc);
        formData.append("task_status", tstatus);
        formData.append("task_start_date", tstart);
        formData.append("task_end_date", tend);

        const fetchPromise = fetch(apiUrl + "/task/create", {
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

                dispatch("changeTodo", "create");

                return response.json();
            })
            .catch((error) => {
                console.log(`新增失败：${error}`);
            });

        backupTodo = { ...todo }; // 更新备份
    }

    function updateTask({
        tname = todo.task_name,
        tdesc = todo.task_desc,
        tstatus = todo.task_status,
        tstart = todo.task_start_date,
        tend = todo.task_end_date,
    }) {
        let formData = new URLSearchParams();
        formData.append("task_name", tname);
        formData.append("task_desc", tdesc);
        formData.append("task_status", tstatus);
        formData.append("task_start_date", tstart);
        formData.append("task_end_date", tend);

        const fetchPromise = fetch(apiUrl + "/task/update/" + todo.task_id, {
            method: "PUT",
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

                dispatch("changeTodo", "update");

                return response.json();
            })
            .catch((error) => {
                console.log(`更新失败：${error}`);
            });

        backupTodo = { ...todo }; // 更新备份
    }

    function deleteTask() {
        if (!confirm("确定删除该任务？")) {
            console.log("取消删除任务");
            return;
        }
        const fetchPromise = fetch(apiUrl + "/task/delete/" + todo.task_id, {
            method: "PUT",
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            },
        });

        fetchPromise
            .then((response) => {
                if (!response.ok) {
                    throw new Error(`HTTP 请求错误：${response.status}`);
                }

                dispatch("changeTodo", "delete");

                return response.json();
            })
            .catch((error) => {
                console.log(`删除失败：${error}`);
            });
    }

    function toggleEdit() {
        isEditing = !isEditing;
    }
    function cancelEdit() {
        todo = backupTodo;
        toggleEdit();
        dispatch("changeTodo", "cancel");
    }
</script>

<li>
    <div class="todo-item-container">
        {#if isEditing}
            <div class="task-container">
                <input
                    class="task-name"
                    type="text"
                    bind:value={todo.task_name}
                />
                <textarea bind:value={todo.task_desc}></textarea>
            </div>
            <div class="task-container">
                <p>Status:&nbsp;</p>
                <select bind:value={todo.task_status}>
                    <option value="active">Active</option>
                    <option value="overdue">Overdue</option>
                    <option value="completed">Completed</option>
                </select>
                <p>
                    Start: <input
                        type="date"
                        bind:value={todo.task_start_date}
                    />
                    => End:
                    <input type="date" bind:value={todo.task_end_date} />
                </p>
            </div>
        {:else}
            <div class="task-container">
                <h3 class="task-name">{todo.task_name}</h3>
                <p>{todo.task_desc}</p>
            </div>
            <div class="task-container">
                <p>Status:&nbsp;</p>
                <p
                    class:status-completed={todo.task_status === "completed"}
                    class:status-overdue={todo.task_status === "overdue"}
                    class:status-default={todo.task_status !== "overdue" &&
                        todo.task_status !== "completed"}
                >
                    {todo.task_status}
                </p>
                <p>
                    Start: {todo.task_start_date} => End: {todo.task_end_date}
                </p>
            </div>
        {/if}
    </div>

    <div class="button-list">
        {#if isEditing}
            <button
                class="button-edit"
                on:click={() => {
                    if (todo.task_id === -1) {
                        createTask({
                            tname: todo.task_name,
                            tdesc: todo.task_desc,
                            tstatus: todo.task_status,
                            tstart: todo.task_start_date,
                            tend: todo.task_end_date,
                        });
                    } else {
                        updateTask({
                            tname: todo.task_name,
                            tdesc: todo.task_desc,
                            tstatus: todo.task_status,
                            tstart: todo.task_start_date,
                            tend: todo.task_end_date,
                        });
                    }

                    toggleEdit();
                }}>Save&Exit</button
            >
            <button class="button-delete" on:click={cancelEdit}
                >Exit Edit</button
            >
        {:else}
            <button
                class="button-complete"
                on:click={() => {
                    updateTask({ tstatus: "completed" });
                }}>Complete!</button
            >
            <button class="button-edit" on:click={toggleEdit}>Edit</button>
            <button class="button-delete" on:click={deleteTask}>Delete</button>
        {/if}
    </div>
</li>

<style>
    li {
        display: flex; /* 设置为Flex容器 */
        align-items: center; /* 子元素在交叉轴上居中对齐 */
        justify-content: space-between; /* 子元素沿主轴分布 */
        border: 2px solid #acacac; /* 设置边框颜色和宽度 */
        padding: 10px; /* 添加内边距以确保内容不会紧贴边框 */
        width: 800px; /* 固定宽度 */
        height: auto; /* 固定高度 */
        box-sizing: border-box; /* 确保边框和内边距的尺寸包含在宽度和高度内 */
        margin-bottom: 1rem; /* 增加底部间距 */
    }
    .todo-item-container {
        display: flex; /* 设置为Flex容器 */
        flex-direction: column; /* 子元素垂直排列 */
        align-items: flex-start; /* 子元素在交叉轴上对齐到起点 */
    }
    .task-container {
        display: flex;
        justify-content: space-between; /* 子元素沿主轴分布 */
        align-items: center; /* 子元素在交叉轴上居中对齐 */
    }
    .task-name {
        margin-right: 1rem;
    }
    .status-completed,
    .status-overdue,
    .status-default {
        margin-right: 1rem;
    }
    .status-completed {
        color: green;
    }
    .status-overdue {
        color: red;
    }
    .button-list {
        display: flex; /* 设置为Flex容器 */
        flex-direction: column; /* 子元素垂直排列 */
        margin-left: auto; /* 推向容器的右侧 */
        align-items: center; /* 子元素在交叉轴上居中对齐 */
    }
    .button-complete,
    .button-edit,
    .button-delete {
        width: 100px; /* 固定宽度 */
        height: 30px; /* 固定高度 */
        margin-bottom: 0.5rem; /* 增加底部间距 */
    }
    .button-complete {
        background-color: #f0f0f0;
    }
    .button-edit {
        background-color: #f0f0f0;
    }
    .button-delete {
        color: #ffecec;
        background-color: red;
    }
</style>
