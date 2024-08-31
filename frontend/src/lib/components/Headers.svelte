<script>
    export let title = "header";
    export let extra = "";
</script>


<div class="flex flex-col bg-cinema-highlight rounded-lg w-full items-center">
    <h2 class="text-cinema-text font-bold text-xl p-4">{title} {#if extra}({extra}){/if}</h2>
</div>