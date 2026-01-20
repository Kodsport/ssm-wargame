export function getAvailableThemes() {
    const themeFiles = [
        'default.css', // lägg till theme filename när du vill lägga till mer :D
        'purple.css',
        'dark.css',
        'kodcentrum.css',
    ];

    return themeFiles.map(file => {
        const filename = file.replace('.css', '');
        const displayName = filename
            .split('-')
            .map(word => word.charAt(0).toUpperCase() + word.slice(1))
            .join(' ');

        return {
            name: displayName,
            filename: filename
        };
    });
}

export async function getThemesFromDirectory() {
    try {
        const themes = import.meta.glob('/public/themes/*.css');
        return Object.keys(themes).map(path => {
            const filename = path.replace('/public/themes/', '').replace('.css', '');
            const displayName = filename
                .split('-')
                .map(word => word.charAt(0).toUpperCase() + word.slice(1))
                .join(' ');

            return {
                name: displayName,
                filename: filename
            };
        });
    } catch (error) {
        return getAvailableThemes();
    }
}
