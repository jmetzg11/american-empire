from django.core.management.base import BaseCommand
from datetime import datetime, timezone, date
from core.models import Event, Tag, Source, Media


class Command(BaseCommand):
    help = 'Seed the database with initial development data'

    def handle(self, *args, **options):
        if Event.objects.count() > 0:
            self.stdout.write(
                self.style.WARNING('Database already seeded')
            )
            return

        active_date = datetime(2025, 6, 15, 0, 0, 0, tzinfo=timezone.utc)

        # Create tags
        tag_names = ["CIA", "Propaganda", "Cold War", "Journalism", "South America", "Human Rights", "Weapons", "Nicaragua", "Iran", "Reagan"]
        tags = {}
        
        for name in tag_names:
            tag, created = Tag.objects.get_or_create(name=name)
            tags[name] = tag

        # Create events with related data
        events_data = [
            {
                'title': 'CIA Operation Mockingbird',
                'date': date(1950, 3, 15),
                'country': 'USA',
                'description': """Operation Mockingbird was a large-scale program of the United States Central Intelligence Agency that began in the early years of the Cold War.

The operation sought to influence media organizations and journalists to disseminate propaganda favorable to CIA interests. This included recruiting leading American journalists into a network to help present the CIA's views, and funded some student and cultural organizations, and magazines as fronts.

Key aspects included:
- Infiltration of major news organizations
- Creation of propaganda materials
- Coordination with foreign intelligence services""",
                'active': active_date,
                'tags': ['CIA', 'Propaganda', 'Cold War', 'Journalism'],
                'sources': [
                    {'name': 'Church Committee Report', 'url': 'https://archive.org/details/churchcommittee'},
                    {'name': 'Declassified CIA Documents', 'url': 'https://cia.gov/mockingbird-files'},
                ],
                'medias': [
                    {'type': 'photo', 'path': '/1/sample.jpg', 'caption': 'Church Committee Report'},
                    {'type': 'photo', 'path': '/1/sample2.jpeg', 'caption': 'Cool Report'},
                    {'type': 'photo', 'path': '/1/sample3.jpeg', 'caption': 'Different Report'},
                    {'type': 'youtube', 'url': 'bDjGJzBdAwY', 'caption': 'Declassified CIA Documents'},
                    {'type': 'youtube', 'url': 'bGAFTaelGRk', 'caption': 'Declassified FBI Documents'},
                ],
            },
            {
                'title': 'Operation Condor South America',
                'date': date(1975, 11, 28),
                'country': 'Chile',
                'description': """Operation Condor was a United States-backed campaign of political repression and state terror involving intelligence operations and assassination.

The operation was officially implemented in November 1975 by the right-wing dictatorships of the Southern Cone of South America. The program was intended to eradicate communist or Soviet influence and ideas, and to suppress active or potential opposition movements against the participating governments.

Coordinated efforts included:
- Cross-border intelligence sharing
- Joint military operations
- Systematic human rights violations""",
                'active': active_date,
                'tags': ['CIA', 'South America', 'Human Rights', 'Cold War'],
                'sources': [
                    {'name': 'National Security Archive', 'url': 'https://nsarchive.gwu.edu/condor-files'},
                    {'name': 'FBI FOIA Release', 'url': 'https://fbi.gov/condor-documents'},
                ],
                'medias': [
                    {'type': 'photo', 'path': '/2/sample.jpeg', 'caption': 'National Security Archive'},
                    {'type': 'youtube', 'url': 'bDjGJzBdAwY', 'caption': 'FBI FOIA Release'},
                ],
            },
            {
                'title': 'Iran-Contra Weapons Sales',
                'date': date(1985, 8, 20),
                'country': 'Nicaragua',
                'description': """The Iran-Contra affair was a political scandal that occurred during the second term of the Reagan Administration.

Senior administration officials secretly facilitated the sale of arms to Iran, which was then under an arms embargo. The officials hoped that the arms sales would secure the release of American hostages and allow U.S. intelligence agencies to fund the Nicaraguan Contras.

This covert operation involved:
- Illegal arms sales to Iran
- Diversion of proceeds to Contra rebels
- Circumvention of congressional oversight""",
                'active': active_date,
                'tags': ['Weapons', 'Nicaragua', 'Iran', 'Reagan'],
                'sources': [
                    {'name': 'Tower Commission Report', 'url': 'https://reagan.library.gov/tower-commission'},
                    {'name': 'Walsh Report', 'url': 'https://justice.gov/walsh-final-report'},
                ],
                'medias': [
                    {'type': 'photo', 'path': '/3/sample.png', 'caption': 'Tower Commission Report'},
                    {'type': 'youtube', 'url': 'bDjGJzBdAwY', 'caption': 'Walsh Report'},
                ],
            },
        ]

        for event_data in events_data:
            # Create event
            event = Event.objects.create(
                title=event_data['title'],
                date=event_data['date'],
                country=event_data['country'],
                description=event_data['description'],
                active=event_data['active'],
            )

            # Add tags
            for tag_name in event_data['tags']:
                event.tags.add(tags[tag_name])

            # Add sources
            for source_data in event_data['sources']:
                Source.objects.create(
                    event=event,
                    name=source_data['name'],
                    url=source_data['url']
                )

            # Add media
            for media_data in event_data['medias']:
                Media.objects.create(
                    event=event,
                    type=media_data['type'],
                    url=media_data.get('url', ''),
                    path=media_data.get('path', ''),
                    caption=media_data['caption']
                )

        self.stdout.write(
            self.style.SUCCESS('Database seeded successfully!')
        )
