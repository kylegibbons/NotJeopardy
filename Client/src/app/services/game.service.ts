import { Injectable } from '@angular/core';
import { Game } from '../models/game';

@Injectable()
export class GameService {

  public game: Game = {
    "id": "1234",
    "categories": [
      {
        "name": "RECENT BESTSELLERS",
        "clues": [
          {
            "answer": "\"Fire & Blood\" is a novel from this author, set centuries before the events in \"Game of Thrones\"",
            "question": "(George) Martin",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "Tara Westover wrote of her quest to be this, the 1-word title of her memoir about leaving her survivalist family to go to college",
            "question": "Educated",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "Both Michael Wolff's \"Fire and Fury\" & Bob Woodward's \"Fear\" mention this building in their subtitles",
            "question": "the White House",
            "dailyDouble": true,
            "answered": false
          },
          {
            "answer": "This bard of Bangor revisited the bestseller list with the creepy entry \"The Outsider\"",
            "question": "Stephen King",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "Marlon Bundo is a type of this animal in books from Mike Pence's daughter & from John Oliver's \"Last Week Tonight\"",
            "question": "a rabbit",
            "dailyDouble": false,
            "answered": false
          }
        ]
      },
      {
        "name": "WOMEN OLYMPIANS",
        "clues": [
          {
            "answer": "American saber specialist Ibtihaj Muhammad wore a hijab competing in this Olympic sport in 2016",
            "question": "fencing",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "In 2016 tennis ace Monica Puig beat 3 Grand Slam champions to win this U.S. territory's first-ever gold medal",
            "question": "Puerto Rico",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "China's Li Xuerui took gold in 2012 in this racquet sport",
            "question": "badminton",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "Jackie Joyner-Kersee twice won gold in this multi-event Olympic contest",
            "question": "the heptathlon",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "Simone Biles won 4 gold medals in Rio: team, all-around, floor exercise & this high-flying event",
            "question": "the vault",
            "dailyDouble": false,
            "answered": false
          }
        ]
      },
      {
        "name": "THE ELEMENTS",
        "clues": [
          {
            "answer": "Alcoa is the nation's largest producer of this element",
            "question": "aluminum",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "Boy, if I had one of this element, atomic No. 28, every time someone said it was first isolated in 1751",
            "question": "nickel",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "Protactinium was once called brevium because the first isotope discovered had only a 70-second one of these",
            "question": "a half-life",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "The first 3 halogen or \"salt-forming\" elements on the periodic table are fluorine, this, bromine",
            "question": "chlorine",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "This element precedes carbonate in a compound used to treat gastric ulcers & is definitely some of your...",
            "question": "bismuth (or calcium)",
            "dailyDouble": false,
            "answered": false
          }
        ]
      },
      {
        "name": "SPELL THE BEANS",
        "comment": "Each response is a type of bean that you're going to have to spell for us.",
        "clues": [
          {
            "answer": "Its name comes from being shaped like an internal organ",
            "question": "K-I-D-N-E-Y",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "A chocolate source--not the spelling Swiss Miss uses",
            "question": "C-A-C-A-O",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "Sizzling cinnamon & juicy pear are 2 popular flavors of it",
            "question": "J-E-L-L-Y",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "Geographic AKA of butter beans",
            "question": "L-I-M-A",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "Chipotle's ingredients include black beans & this other 5-letter type",
            "question": "P-I-N-T-O",
            "dailyDouble": false,
            "answered": false
          }
        ]
      },
      {
        "name": "WORDS IN TEACHER",
        "clues": [
          {
            "answer": "Third from the Sun",
            "question": "Earth",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "To provide food for an event",
            "question": "cater",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "To deceive for profit",
            "question": "cheat",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "It was originally reckoned as the amount of land a team of oxen could plow in a day",
            "question": "an acre",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "Jasper & agate are 2 of the many minerals classified as this form of quartz",
            "question": "chert",
            "dailyDouble": false,
            "answered": false
          }
        ]
      },
      {
        "name": "EDUCATIONAL GAMES",
        "clues": [
          {
            "answer": "This \"amphibious\" company makes the educational line of LeapPad tablets",
            "question": "LeapFrog",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "In a geography computer game, you need to find out \"Where in the World Is\" she",
            "question": "Carmen Sandiego",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "The Amazon Trail is a spin-off from this classic history game",
            "question": "The Oregon Trail",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "Ms. Frizzle & the yellow title transport feature in web games based on this TV show",
            "question": "The Magic School Bus",
            "dailyDouble": false,
            "answered": false
          },
          {
            "answer": "Mavis Beacon, who has taught this life skill to millions, is not a real person",
            "question": "typing",
            "dailyDouble": false,
            "answered": false
          }
        ]
      }
    ]
  };
  

  constructor() { }

}
