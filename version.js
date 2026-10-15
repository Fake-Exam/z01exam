(() => {
    const badge = document.getElementById('version-badge');
    fetch('/api/version', { cache: 'no-store' })
        .then(response => {
            if (!response.ok) throw new Error('Version unavailable');
            return response.json();
        })
        .then(({ sha, dirty }) => {
            if (!/^[a-f0-9]{40}$/.test(sha)) throw new Error('Unknown revision');
            badge.textContent = `${sha.slice(0, 8)}${dirty ? ' · modifié' : ''}`;
            badge.title = `Version exécutée : ${sha}${dirty ? ' (modifications locales)' : ''}`;
        })
        .catch(() => { badge.textContent = 'version inconnue'; });
})();
