<script>
    import '../styles/global.scss';
    import Header from '$lib/components/Header.svelte';
</script>

<svelte:head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</svelte:head>

<main class="bg-cinema-primary h-svh antialiased">
    <Header />
    <slot></slot>
</main>
