<script>
    import '../styles/global.scss';
    import Header from '$lib/components/Header.svelte';
</script>


<main class="bg-cinema-primary h-svh antialiased">
    <Header />
    <slot></slot>
</main>
