// Search Modal functionality
document.addEventListener('DOMContentLoaded', () => {
    const searchBtn = document.getElementById('searchBtn');
    const searchModal = document.getElementById('searchModal');
    const closeModal = document.getElementById('closeModal');
    const tagSearchInput = document.getElementById('tagSearchInput');
    const tagResults = document.getElementById('tagResults');

    let allTags = [];

    if (searchBtn && searchModal && closeModal) {
        searchBtn.addEventListener('click', async () => {
            searchModal.classList.remove('hidden');

            // Fetch tags if not already loaded
            if (allTags.length === 0) {
                await fetchTags();
            }

            // Focus on search input
            if (tagSearchInput) {
                tagSearchInput.focus();
            }
        });

        closeModal.addEventListener('click', () => {
            searchModal.classList.add('hidden');
            if (tagSearchInput) {
                tagSearchInput.value = '';
            }
        });

        // Close modal when clicking outside
        searchModal.addEventListener('click', (e) => {
            if (e.target === searchModal) {
                searchModal.classList.add('hidden');
                if (tagSearchInput) {
                    tagSearchInput.value = '';
                }
            }
        });

        // Close modal with Escape key
        document.addEventListener('keydown', (e) => {
            if (e.key === 'Escape' && !searchModal.classList.contains('hidden')) {
                searchModal.classList.add('hidden');
                if (tagSearchInput) {
                    tagSearchInput.value = '';
                }
            }
        });

        // Search input filtering
        if (tagSearchInput) {
            tagSearchInput.addEventListener('input', (e) => {
                const searchTerm = e.target.value.toLowerCase();
                const filteredTags = allTags.filter(tag =>
                    tag.toLowerCase().includes(searchTerm)
                );
                displayTags(filteredTags);
            });
        }
    }

    async function fetchTags() {
        try {
            const response = await fetch('/search');
            if (!response.ok) {
                throw new Error('Failed to fetch tags');
            }
            allTags = await response.json();
            displayTags(allTags);
        } catch (error) {
            console.error('Error fetching tags:', error);
            if (tagResults) {
                tagResults.innerHTML = '<div class="text-red-500 text-center py-4">Error loading tags</div>';
            }
        }
    }

    function displayTags(tags) {
        if (!tagResults) return;

        if (tags.length === 0) {
            tagResults.innerHTML = '<div class="text-gray-500 text-center py-4">No tags found</div>';
            return;
        }

        const tagsHTML = '<div class="flex flex-wrap gap-2 p-2">' + tags.map(tag => `
            <a href="/?tags=${encodeURIComponent(tag)}"
               class="bg-blue-200 rounded-xl py-1 px-2.5 inline-block hover:bg-blue-300 transition-colors">
                ${escapeHtml(tag)}
            </a>
        `).join('') + '</div>';

        tagResults.innerHTML = tagsHTML;
    }

    function escapeHtml(text) {
        const div = document.createElement('div');
        div.textContent = text;
        return div.innerHTML;
    }

    // Books expandable rows functionality (table view)
    const bookRows = document.querySelectorAll('.book-row');
    if (bookRows.length > 0) {
        bookRows.forEach(row => {
            row.addEventListener('click', function() {
                const index = this.getAttribute('data-book-index');
                const eventsRow = document.getElementById('events-' + index);
                const toggleIcon = this.querySelector('.toggle-icon');

                if (eventsRow.classList.contains('hidden')) {
                    eventsRow.classList.remove('hidden');
                    toggleIcon.textContent = '▼';
                } else {
                    eventsRow.classList.add('hidden');
                    toggleIcon.textContent = '▶';
                }
            });
        });
    }

    // Books expandable cards functionality (mobile view)
    const bookCards = document.querySelectorAll('.book-card');
    if (bookCards.length > 0) {
        bookCards.forEach(card => {
            card.addEventListener('click', function() {
                const index = this.getAttribute('data-book-index');
                const eventsCard = document.getElementById('events-card-' + index);
                const toggleIcon = this.querySelector('.toggle-icon');

                if (eventsCard.classList.contains('hidden')) {
                    eventsCard.classList.remove('hidden');
                    toggleIcon.textContent = '▼';
                } else {
                    eventsCard.classList.add('hidden');
                    toggleIcon.textContent = '▶';
                }
            });
        });
    }
});
