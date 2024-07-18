<script>
    export let title = "Movie";
    export let poster = "";
    export let link = "#";
</script>



<div class="flex flex-col shrink-0 w-1/2 rounded-lg p-2">
    <a href={link} class="flex">
        <img class="object-cover rounded-lg w-full h-auto" src={poster} alt={title} />
    </a>
    <a href={link} class="text-cinema-text line-clamp-2 py-4 ">{title}</a>
</div>