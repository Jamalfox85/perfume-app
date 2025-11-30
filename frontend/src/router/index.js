import { createRouter, createWebHistory } from "vue-router";
import { supabase } from "@/auth/supabase";

import Home from "../views/Home.vue";
import Perfumes from "@/views/Perfumes.vue";
import Brands from "@/views/Brands.vue";
import Notes from "@/views/Notes.vue";
import Collections from "@/views/Collections.vue";
import MyCabinet from "@/views/MyCabinet.vue";

import Profile from "@/views/Profile.vue";
import Settings from "@/views/Settings.vue";

import Login from "@/views/Login.vue";
import AuthCallback from "@/views/AuthCallback.vue";
import Signup from "@/views/Signup.vue";

const routes = [
    {
        path: "/",
        name: "Home",
        component: Home,
        meta: {
            requiresAuth: false,
            nav: true,
            label: "Home",
            icon: "fa-solid fa-house",
        },
    },
    {
        path: "/perfumes",
        name: "Perfumes",
        component: Perfumes,
        meta: {
            requiresAuth: false,
            nav: true,
            label: "Perfumes",
            icon: "fa-solid fa-bottle-droplet",
        },
    },
    {
        path: "/brands",
        name: "Brands",
        component: Brands,
        meta: {
            requiresAuth: false,
            nav: true,
            label: "Brands",
            icon: "fa-solid fa-building",
        },
    },
    {
        path: "/notes",
        name: "Notes",
        component: Notes,
        meta: {
            requiresAuth: false,
            nav: true,
            label: "Notes",
            icon: "fa-brands fa-pagelines",
        },
    },
    {
        path: "/collections",
        name: "Collections",
        component: Collections,
        meta: {
            requiresAuth: false,
            nav: true,
            label: "Collections",
            icon: "fa-solid fa-table-list",
        },
    },
    {
        path: "/my-cabinet",
        name: "MyCabinet",
        component: MyCabinet,
        meta: {
            requiresAuth: true,
            nav: true,
            label: "My Cabinet",
            icon: "fa-solid fa-suitcase-rolling",
        },
    },
    {
        path: "/profile",
        name: "Profile",
        component: Profile,
        meta: { requiresAuth: true, nav: false },
    },
    {
        path: "/settings",
        name: "Settings",
        component: Settings,
        meta: { requiresAuth: true, nav: false },
    },
    {
        path: "/login",
        name: "Login",
        component: Login,
        meta: { nav: false },
    },
    {
        path: "/auth/callback",
        name: "AuthCallback",
        component: AuthCallback,
        meta: { nav: false },
    },
    {
        path: "/signup",
        name: "Signup",
        component: Signup,
        meta: { nav: false },
    },
];

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
});

router.beforeEach(async (to, from, next) => {
    const { data } = await supabase.auth.getSession();

    const isLoggedIn = !!data.session;

    if (to.meta.requiresAuth && !isLoggedIn) {
        return next("/login");
    }

    next();
});

export default router;
