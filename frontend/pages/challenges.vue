<template>
    <div class="challenges-container">
        <div class="header-section mb-3">
            <h1 class="text-primary h3 mb-2">Utmaningar</h1>
            <p class="mb-0 small">
                Här kan du öva på utmaningar från tidigare års upplagor av Säkerhets-SM. 
            </p>
        </div>
        <div class="filter-bar mb-3">
            <div class="row g-2">
                <div class="col-md-3">
                    <label class="form-label extra-small">Sök</label>
                    <input 
                        type="text" 
                        class="form-control form-control-sm" 
                        placeholder="Sök utmaningar..." 
                        v-model="searchQuery"
                    >
                </div>
                <div class="col-md-2">
                    <label class="form-label extra-small">Kategori</label>
                    <select class="form-select form-select-sm" v-model="challFilter.categoryFilter">
                        <option value="">Alla kategorier </option>
                        <option v-for="cat in categories" :key="cat" :value="cat">
                            {{ cat }} ({{ getCategoryCount(cat) }})
                        </option>
                    </select>
                </div>
<!--                 <div class="col-md-2">
                    <label class="form-label extra-small">Svårighetsgrad</label>
                    <select class="form-select form-select-sm" v-model="difficultyFilter">
                        <option value="">Alla nivåer</option>
                        <option value="enkel">Enkel ({{ getDifficultyCount('enkel') }})</option>
                        <option value="medel">Medel ({{ getDifficultyCount('medel') }})</option>
                        <option value="svår">Svår ({{ getDifficultyCount('svår') }})</option>
                    </select>
                </div> -->
                <div class="col-md-2">
                    <label class="form-label extra-small">Status</label>
                    <select class="form-select form-select-sm" v-model="statusFilter">
                        <option value="">Alla</option>
                        <option value="solved">Löst</option>
                        <option value="unsolved">Ej löst</option>
                    </select>
                </div>
                <div class="col-md-2">
                    <label class="form-label extra-small">Sortera</label>
                    <div class="d-flex gap-1">
                        <select class="form-select form-select-sm flex-grow-1" v-model="sortBy">
<!--                             <option value="difficulty">Svårighet</option> -->
                            <option value="title">Titel</option>
                            <option value="score">Poäng</option>
                            <option value="solves">Antal lösningar</option>
                        </select>
                        <button 
                            class="btn btn-secondary btn-sm" 
                            @click="toggleSortOrder"
                            :title="sortOrder === 'asc' ? 'Sortera stigande' : 'Sortera fallande'"
                        >
                            <i class="bi" :class="sortOrder === 'asc' ? 'bi-sort-up' : 'bi-sort-down'"></i>
                        </button>
                    </div>
                </div>
                <div class="col-md-1 d-flex align-items-end">
                    <button class="btn btn-secondary btn-sm" @click="clearFilters">
                        Rensa
                    </button>
                </div>
            </div>
        </div>

        <div class="event-filter mb-3" v-if="challs.events.length">
            <button 
                class="btn btn-outline-primary btn-sm mb-1" 
                @click="showEventFilter = !showEventFilter"
                type="button"
            >
                <i class="bi bi-filter"></i>
                Filtrera på tävling ({{ selectedEventsCount }} valda)
            </button>
            <div class="collapse" :class="{ show: showEventFilter }">
                <div class="card card-body py-2">
                    <div class="row g-1">
                        <div class="col-auto" v-for="event in (challs.events as any[])" :key="event.id">
                            <div class="form-check">
                                <input 
                                    type="checkbox" 
                                    class="form-check-input" 
                                    :id="'event-' + event.id"
                                    v-model="(challFilter.eventFilter as any)[event.id]"
                                >
                                <label class="form-check-label extra-small" :for="'event-' + event.id">
                                    {{ event.name }}
                                </label>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <div class="view-controls mb-2 d-flex justify-content-between align-items-center">
            <div class="btn-group btn-group-sm" role="group">
                <input type="radio" class="btn-check" id="grid-view" v-model="viewMode" value="grid">
                <label class="btn btn-primary me-2" for="grid-view">
                    <i class="bi bi-grid-3x3-gap"></i> Rutnät
                </label>
                <input type="radio" class="btn-check" id="list-view" v-model="viewMode" value="list">
                <label class="btn btn-primary" for="list-view">
                    <i class="bi bi-list"></i> Lista
                </label>
            </div>
            <div v-if="viewMode === 'grid'" class="d-flex align-items-center gap-2">
                <label class="form-label extra-small mb-0">Utmaningar per rad:</label>
                <input 
                    type="number" 
                    class="form-control form-control-sm" 
                    style="width: 80px;" 
                    v-model.number="challengesPerRow"
                    min="1"
                    max="50"
                    step="1"
                >
            </div>
        </div>
        <div class="grid-divider mb-3"></div>

        <div v-if="viewMode === 'grid'" class="ssm-grid">
            <template v-for="(chall, index) in displayedChallenges" :key="chall.id">
                <div>
                    <ChallengePreview 
                        class="pointer challenge-compact" 
                        @click="nav(chall.slug)" 
                        :chall="chall" 
                    />
                    <a class="d-none" :href="`/challenges/${chall.slug}`">{{ chall.title }}</a>
                </div>
                <div v-if="(index + 1) % challengesPerRow === 0 && index < displayedChallenges.length - 1" class="grid-divider"></div>
            </template>
        </div>

        <div v-else class="list-view">
            <div class="table-responsive">
                <table class="table table-hover table-sm">
                    <thead class="table">
                        <tr>
                            <th class="extra-small">Titel</th>
                            <th class="extra-small">Kategori</th>
                            <!-- <th class="extra-small">Svårighetsgrad</th> -->
                            <th class="extra-small">Poäng</th>
                            <th class="extra-small">Lösningar</th>
                            <th class="extra-small">Status</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="chall in displayedChallenges" :key="chall.id" class="pointer" @click="nav(chall.slug)">
                            <td>
                                <strong class="small">{{ chall.title }}</strong>
                                <div class="extra-small text">{{ chall.description?.substring(0, 40) }}...</div>
                            </td>
                            <td>
                                <span class="badge extra-small" :class="getCategoryClass(chall)">{{ chall.category }}</span>
                            </td>
                            <td class="small">{{ chall.score || chall.static_score }}</td>
                            <td class="small">{{ chall.solves || 0 }}</td>
                            <td>
                                <span v-if="chall.solved" class="badge bg-success extra-small">Löst</span>
                                <span v-else class="badge bg-warning extra-small">Ej löst</span>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>

        <div v-if="isLoading" class="text-center py-3">
            <div class="spinner-border spinner-border-sm text-primary" role="status">
                <span class="visually-hidden">Loading...</span>
            </div>
            <div class="small text mt-2">Laddar fler utmaningar...</div>
        </div>

        <div v-if="canLoadMore && !isLoading" class="text-center py-3">
            <button class="btn btn-outline-primary btn-sm" @click="loadMore">
                Ladda fler utmaningar
            </button>
        </div>

        <div v-if="!canLoadMore && displayedCount > 0 && filteredChallenges.length > itemsPerLoad" class="text-center py-3">
            <small class="text">Alla utmaningar har laddats</small>
        </div>

        <div v-if="filteredChallenges.length === 0" class="text-center py-4">
            <h6 class="text">Inga utmaningar hittades</h6>
            <p class="text small">Försök ändra dina filterinställningar</p>
            <button class="btn btn-primary btn-sm" @click="clearFilters">Rensa alla filter</button>
        </div>
    </div>
    <NuxtPage />
</template>

<script setup lang="ts">
import { useChallengeStore } from '../store/challenges'
import { storeToRefs } from "pinia";

useHead({
    title: 'SSM - Utmaningar'
})

useServerSeoMeta({
    title: 'SSM - Utmaningar'
})

const challs = useChallengeStore()
const router = useRouter()

const { challFilter } = storeToRefs(useChallengeStore())

const searchQuery = ref('')
const difficultyFilter = ref('')
const statusFilter = ref('')
const sortBy = ref('title')
const sortOrder = ref('asc')
const viewMode = ref('grid')
const showEventFilter = ref(false)
const challengesPerRow = ref(3)

// Infinite scroll variables
const displayedCount = ref(20)
const itemsPerLoad = 20
const isLoading = ref(false)

await useAsyncData('challenges', challs.getChallenges)
await useAsyncData('events', challs.getEvents)

onMounted(() => {
    challs.getChallenges()
})

const categories = computed(() => 
    (challs.challenges as any[])
        .map((c: any) => c.category)
        .filter((v: any, i: number, a: any[]) => a.indexOf(v) === i)
        .sort()
)

/* const getDifficulty = (chall: any) => {
    const score = chall.score || 0
    if (score <= 150) return 'enkel'
    if (score <= 300) return 'medel'
    if (score <= 6000) return 'svår'
}

const getDifficultyClass = (chall: any) => {
    const difficulty = getDifficulty(chall)
    return {
        'difficulty-easy': difficulty === 'enkel',
        'difficulty-medium': difficulty === 'medel',
        'difficulty-hard': difficulty === 'svår',
    }
} */

const filteredChallenges = computed(() => {
    let result = (challs.challenges as any[]).filter((c: any) => c.chall_namespace == null)
    
    const allowedEvents: string[] = []
    for (const [key, value] of Object.entries(challFilter.value.eventFilter)) {
        if (value) allowedEvents.push(key)
    }
    if (allowedEvents.length > 0) {
        result = result.filter((c: any) => allowedEvents.includes(c.ctf_event_id))
    }
    
    if (challFilter.value.categoryFilter) {
        result = result.filter((c: any) => c.category === challFilter.value.categoryFilter)
    }
    
    if (searchQuery.value) {
        const query = searchQuery.value.toLowerCase()
        result = result.filter((c: any) => 
            c.title.toLowerCase().includes(query) ||
            c.description?.toLowerCase().includes(query) ||
            c.category.toLowerCase().includes(query)
        )
    }
    
/*     if (difficultyFilter.value) {
        result = result.filter((c: any) => getDifficulty(c) === difficultyFilter.value)
    } */
    
    if (statusFilter.value === 'solved') {
        result = result.filter((c: any) => c.solved)
    } else if (statusFilter.value === 'unsolved') {
        result = result.filter((c: any) => !c.solved)
    }
    
    return result
})

const sortedChallenges = computed(() => {
    const sorted = [...filteredChallenges.value]
    const isAsc = sortOrder.value === 'asc'
    
    switch (sortBy.value) {
        case 'title':
            return sorted.sort((a: any, b: any) => {
                const result = a.title.localeCompare(b.title)
                return isAsc ? result : -result
            })
        case 'score':
            return sorted.sort((a: any, b: any) => {
                const result = (a.score || a.static_score || 0) - (b.score || b.static_score || 0)
                return isAsc ? result : -result
            })
        case 'solves':
            return sorted.sort((a: any, b: any) => {
                const result = (a.solves || 0) - (b.solves || 0)
                return isAsc ? result : -result
            })
        case 'newest':
            return sorted.sort((a: any, b: any) => {
                const result = new Date(a.created_at || 0).getTime() - new Date(b.created_at || 0).getTime()
                return isAsc ? result : -result
            })
        case 'difficulty':
        default:
            /* const difficultyOrder: { [key: string]: number } = { 'enkel': 1, 'medel': 2, 'svår': 3 }
            return sorted.sort((a: any, b: any) => {
                const diffA = getDifficulty(a) || 'svår'
                const diffB = getDifficulty(b) || 'svår'
                const orderA = difficultyOrder[diffA] || 4
                const orderB = difficultyOrder[diffB] || 4
                if (orderA !== orderB) {
                    const result = orderA - orderB
                    return isAsc ? result : -result
                }
                const scoreResult = (a.score || a.static_score || 0) - (b.score || b.static_score || 0)
                return isAsc ? scoreResult : -scoreResult
            }) */
            return sorted.sort((a: any, b: any) => {
                const result = a.title.localeCompare(b.title)
                return isAsc ? result : -result
            })
    }
})

const displayedChallenges = computed(() => {
    return sortedChallenges.value.slice(0, displayedCount.value)
})

const canLoadMore = computed(() => {
    return displayedCount.value < sortedChallenges.value.length
})

const loadMore = () => {
    if (canLoadMore.value && !isLoading.value) {
        isLoading.value = true
        setTimeout(() => {
            displayedCount.value += itemsPerLoad
            isLoading.value = false
        }, 80)
    }
}

const handleScroll = () => {
    if (isLoading.value || !canLoadMore.value) return
    
    const scrollTop = window.pageYOffset || document.documentElement.scrollTop
    const windowHeight = window.innerHeight
    const documentHeight = document.documentElement.scrollHeight
    
    if (scrollTop + windowHeight >= documentHeight - 200) {
        loadMore()
    }
}

onMounted(() => {
    window.addEventListener('scroll', handleScroll)
})

onUnmounted(() => {
    window.removeEventListener('scroll', handleScroll)
})

const getCategoryCount = (category: string) => {
    return (challs.challenges as any[]).filter((c: any) => c.category === category && c.chall_namespace == null).length
}

/* const getDifficultyCount = (difficulty: string) => {
    return (challs.challenges as any[]).filter((c: any) => getDifficulty(c) === difficulty && c.chall_namespace == null).length
} */

const getCategoryClass = (chall: any) => {
    const category = chall.category?.toLowerCase()
    switch (category) {
        case 'pwn':
            return 'category-pwn'
        case 'crypto':
            return 'category-crypto'
        case 'web':
            return 'category-web'
        case 'forensics':
            return 'category-forensics'
        case 'misc':
            return 'category-misc'
        case 'osint':
            return 'category-osint'
        case 'rev':
        case 'reverse':
        case 'reversing':
            return 'category-rev'
        default:
            return 'bg-secondary'
    }
}

const selectedEventsCount = computed(() => {
    return Object.values(challFilter.value.eventFilter).filter(Boolean).length
})

const clearFilters = () => {
    challFilter.value.categoryFilter = ''
    ;(challFilter.value.eventFilter as any) = {}
    searchQuery.value = ''
    difficultyFilter.value = ''
    statusFilter.value = ''
    displayedCount.value = itemsPerLoad // reset till default
}

const toggleSortOrder = () => {
    sortOrder.value = sortOrder.value === 'asc' ? 'desc' : 'asc'
}

function nav(slug: string) {
    router.push(`/challenges/${slug}`)
}

function selectChallenge(challenge: any) {
    router.push(`/challenges/${challenge.slug}`)
}

watch([searchQuery, difficultyFilter, statusFilter, sortBy, sortOrder, () => challFilter.value.categoryFilter], () => {
    displayedCount.value = itemsPerLoad
})
</script>

<style scoped>
.challenges-container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 2rem;
}

.filter-bar {
    background: #001125;
    border-radius: 6px;
    padding: 0.75rem;
    border: 1px solid #ffab1c;
}

.form-label {
    font-weight: 500;
    margin-bottom: 0.15rem;
}

.extra-small {
    font-size: 0.7rem;
}

.pointer {
    cursor: pointer;
}

.ssm-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
    gap: 0.75rem;
    max-width: 100%;
    margin: 0 auto;
}

.challenge-compact {
    transform: scale(1);
    transform-origin: center;
    transition: transform 0.2s ease-in-out;
}

.challenge-compact:hover {
    transform: scale(1.02);
}

.grid-divider {
    grid-column: 1 / -1;
    height: 2px;
    background-color: #ffab1c;
    margin: 1rem 0;
    border-radius: 1px;
}

.list-view .table th {
    border-top: none;
    font-weight: 600;
    color: #ffffff;
    padding: 0.5rem 0.25rem;
}

.list-view .table td {
    vertical-align: middle;
    padding: 0.5rem 0.25rem;
}

.difficulty-badge {
    padding: 0.15rem 0.35rem;
    border-radius: 0.2rem;
    font-size: 0.65rem;
    font-weight: 600;
    border: 1px solid #00000077 !important;
}

.difficulty-easy {
    background-color: #d4edda;
    color: #155724;
}

.difficulty-medium {
    background-color: #fff3cd;
    color: #856404;
}

.difficulty-hard {
    background-color: #f8d7da;
    color: #721c24;
}

.difficulty-unsolved {
    background-color: #e2e3e5;
    color: #383d41;
}

.view-controls {
    padding-bottom: 0.5rem;
}

.btn-check:checked + .btn {
    background-color: var(--bs-primary);
    border-color: #d97100;
}

.event-filter .card {
    border: 1px solid #dee2e6;
}

.event-filter .form-check-input:checked {
    background-color: #d97100 !important;
    border-color: #d97100 !important;
}

.event-filter .form-check-input:focus {
    border-color: #d97100 !important;
    box-shadow: 0 0 0 0.25rem rgba(217, 113, 0, 0.25) !important;
}

.header-section h1 {
    margin-bottom: 0.25rem;
}

.header-section p {
    margin-bottom: 0;
    font-size: 0.95rem;
}

.table-responsive {
    border-radius: 0.25rem;
    overflow: hidden;
}

.pagination {
    margin-bottom: 0;
}

.pagination .page-link {
    color: var(--bs-primary);
    border-color: #dee2e6;
    padding: 0.25rem 0.5rem;
}

.pagination .page-item.active .page-link {
    background-color: var(--bs-primary);
    border-color: #d97100;
}

.pagination .page-item.disabled .page-link {
    color: #6c757d;
    background-color: #fff;
    border-color: #dee2e6;
}

.badge {
    font-size: 0.65rem;
    padding: 0.25rem 0.4rem;
}

.badge.bg-warning {
    background-color: #ffbc11f0 !important;
    color: white !important;
}

.category-pwn {
    background-color: #b01d2b;
    color: white;
}

.category-crypto {
    background-color: #ffbf00;
    color: #ffffff;
}

.category-web {
    background-color: #238cb6;
    color: #ffffff;
}

.category-forensics {
    background-color: #144690;
    color: white;
}

.category-misc {
    background-color: #6f42c1;
    color: white;
}

.category-osint {
    background-color: #6c757d;
    color: white;
}

.category-rev {
    background-color: #b47523;
    color: white;
}

@media (max-width: 1200px) {
    .challenges-container {
        padding: 0 1.5rem;
    }
    
    .ssm-grid {
        grid-template-columns: repeat(3, 1fr);
        gap: 0.9rem;
    }
}

@media (max-width: 768px) {
    .challenges-container {
        padding: 0 1rem;
    }
    
    .ssm-grid {
        grid-template-columns: repeat(2, 1fr);
        gap: 0.8rem;
    }
    
    .challenge-compact {
        transform: scale(1);
    }
    
    .challenge-compact:hover {
        transform: scale(1.02);
    }
    
    .filter-bar .row {
        row-gap: 0.75rem;
    }
    
    .view-controls {
        flex-direction: column;
        gap: 0.25rem;
        align-items: flex-start !important;
    }
}

@media (max-width: 576px) {
    .challenges-container {
        padding: 0 0.75rem;
    }
    
    .ssm-grid {
        grid-template-columns: 1fr;
        gap: 0.6rem;
    }
    
    .filter-bar {
        padding: 0.5rem;
    }
    
    .table-responsive {
        font-size: 0.8rem;
    }
    
    .challenge-compact {
        transform: scale(1);
    }
    
    .challenge-compact:hover {
        transform: scale(1.02);
    }
}

@media (max-width: 400px) {
    .challenges-container {
        padding: 0 0.5rem;
    }
}

.list-view .table tbody tr:hover {
    background-color: rgba(0, 123, 255, 0.05);
}

.form-control-sm {
    padding: 0.25rem 0.5rem;
    font-size: 0.8rem;
}

.form-select-sm {
    padding: 0.25rem 0.5rem;
    font-size: 0.8rem;
}

.btn-sm {
    padding: 0.25rem 0.5rem;
    font-size: 0.8rem;
}

.loading-skeleton {
    background: linear-gradient(90deg, #f0f0f0, #e0e0e0, #f0f0f0);
    background-size: 200% 100%;
    animation: loading 1.5s infinite;
}

@keyframes loading {
    0% {
        background-position: 200% 0;
    }
    100% {
        background-position: -200% 0;
    }
}
</style>
