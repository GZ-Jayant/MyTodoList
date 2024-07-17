import { writable } from 'svelte/store';

export const config = writable({}); // 初始化一个空对象

// 异步加载配置数据并更新store
try {
    fetch('/config.json')
        .then(response => response.json())
        .then(data => {
            config.set(data);
            console.log("加载配置文件成功:", data);
        });
} catch (error) {
    console.error("加载配置文件失败:", error);
}
