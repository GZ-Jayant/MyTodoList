import App from './App.svelte';
import { config } from "./configStore.js"; // 导入配置加载函数


let app;
// 订阅配置数据的变化
config.subscribe(($config) => {
	// 检查配置数据是否已加载
	if ($config.isLoaded) {
		// 配置数据加载完成，初始化应用
		app = new App({
			target: document.body,
			props: {
				apiUrl: $config.apiBaseUrl // 将apiUrl作为prop传递
			}
		});
	}
});

export default app;