import { watch, onMounted, onUnmounted, computed, type Ref } from 'vue'

export function useTheme(themeRef: Ref<string> | (() => string) | null) {
  let currentThemeLink: HTMLLinkElement | null = null

  function loadTheme(themeName?: string) {
    if (currentThemeLink) {
      currentThemeLink.remove()
      currentThemeLink = null
    }

    const theme = themeName || (typeof themeRef === 'function' ? themeRef() : themeRef?.value) || 'ctf-theme'


    const link = document.createElement('link')
    link.rel = 'stylesheet'
    link.href = `/themes/${theme}.css`
    link.setAttribute('data-ctf-theme', 'true')

    link.onerror = () => {

      const fallbackLink = document.createElement('link')
      fallbackLink.rel = 'stylesheet'
      fallbackLink.href = `/assets/themes/${theme}.css`
      fallbackLink.setAttribute('data-ctf-theme', 'true')

      fallbackLink.onload = () => {
        currentThemeLink = fallbackLink
      }

      fallbackLink.onerror = () => {
        if (theme !== 'ctf-theme') {
          loadTheme('ctf-theme')
        }
      }

      document.head.appendChild(fallbackLink)
      return
    }

    document.head.appendChild(link)
    currentThemeLink = link
  }

  function cleanupTheme() {
    if (currentThemeLink) {
      currentThemeLink.remove()
      currentThemeLink = null
    }
  }

  if (themeRef && typeof themeRef === 'object' && 'value' in themeRef) {
    const stopWatcher = watch(themeRef, (newTheme: string) => {
      if (newTheme) {
        loadTheme(newTheme)
      }
    }, { immediate: false })

    onMounted(() => {
      loadTheme()
    })

    onUnmounted(() => {
      cleanupTheme()
      stopWatcher()
    })
  }

  return {
    loadTheme,
    cleanupTheme
  }
}

export function useCTFTheme(ctfStore: any) {
  const ctfTheme = computed(() => ctfStore?.ctf?.theme || 'ctf-theme');
  return useTheme(ctfTheme);
}

export function useManualTheme() {
  return useTheme(null)
}
