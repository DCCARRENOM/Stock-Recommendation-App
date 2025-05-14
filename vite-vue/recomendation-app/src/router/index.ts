import HelloWorld from "../components/Home.vue";
import { createRouter, createWebHistory } from "vue-router";

const routes = [
    {path: '/home', component: HelloWorld},
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;