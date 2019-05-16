import { getPluralCategory } from "@angular/common/src/i18n/localization";
import { clearModulesForTest } from "@angular/core/src/linker/ng_module_factory_loader";

export interface Game {
    id: string;
    creatorId?: string;
    creatorName?: string;
    gameName?: string;
    rounds: Round[];
    round: number;
    activeCategory: Category;
    activeClue: Clue;
}

export interface Round {
    name: string;
    comment?: string;
    categories: Category[];
}
    
export interface Category {
    name: string;
    comment?: string;
    clues: Clue[];
    media?: string;
}

export interface Clue {
    comment?: string;
    answered: boolean;
    answer: string;
    question: string;
    media?: string;
    dailyDouble: boolean;
}

export interface ClueSelect {
    round: number;
    categoryNumber: number;
    clueNumber: number;
}