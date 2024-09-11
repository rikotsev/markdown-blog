export default class ArticleApi {
    private static instance: ArticleApi;

    public static getInstance(): ArticleApi {
        if (!ArticleApi.instance) {
            ArticleApi.instance = new ArticleApi()
        }

        return ArticleApi.instance
    }

    private categories: Category[];

    private articles: Map<string, Article[]> = new Map<string, Article[]>()

    private constructor() {
        this.categories = [
            {
                id: "cat_1",
                prettyId: "tech",
                title: "Tech"
            }, {
                id: "cat_2",
                prettyId: "philosophy",
                title: "Philosophy"
            },
            {
                id: "cat_3",
                prettyId: "travel",
                title: "Travel"
            }
        ]

        this.articles.set(this.categories[0].prettyId, [
            {
                id: "article_1",
                prettyId: "new_js_framework",
                title: "New JS Framework",
                description: "Yet again - a new JS framework claims it can do everything",
                content: "# New JS Framework \n" +
                "--- \n" +
                "Think about all the possibilities only with this new JS framework\n" +
                "* [X] it's new\n" +
                "* [X] it's javascript\n" +
                "* [ ] it's useful\n" +
                "\n" +
                "how is this so popular in today's world honestly?"
            },
            {
                id: "article_2",
                prettyId: "writing_a_db_from_scratch",
                title: "Writing a DB from scratch",
                description: "My trying to implement some BTrees in C",
                content: "# Writing a DB from scratch"
            },
            {
                id: "article_3",
                prettyId: "http_server_from_scratch",
                title: "HTTP Server from scratch",
                description: "Doing a simple http server in C",
                content: "# HTTP Server from scratch"
            }
        ]);

        this.articles.set(this.categories[1].prettyId, [
            {
                id: "article_4",
                prettyId: "socrates_and_his_dialogues",
                title: "Socrates and his dialogues",
                description: "Look at the 5 most interesting dialogues",
                content: "# Socrates and his dialogues"
            },
            {
                id: "article_5",
                prettyId: "kant_and_his_logic",
                title: "Kant and his logic",
                description: "Let's take a look at critique of the rational mind...",
                content: "# Kant and his logic"
            },
            {
                id: "article_6",
                prettyId: "motorcycle_maintenance_and_me",
                title: "Motorcycle Maintenance and Me",
                description: "How that book impacted me",
                content: "# Motorcycle Maintenance and Me"
            }
        ]);

        this.articles.set(this.categories[2].prettyId, [
            {
                id: "article_7",
                prettyId: "granada",
                title: "Granada",
                description: "Granada - mix of many cultures",
                content: "# Granada"
            },
            {
                id: "article_8",
                prettyId: "matera",
                title: "Matera",
                description: "Matera - the shameful past",
                content: "# Matera"
            },
            {
                id: "article_9",
                prettyId: "rupite",
                title: "Rupite",
                description: "Rupite - mysticism and more",
                content: "# Rupite"
            }
        ]);

    }

    public getArticles(categoryPrettyId: string): Article[] {
        const result = this.articles.get(categoryPrettyId)
        if (!result) {
            return []
        }

        return result
    }

    public getAllArticles(): Article[] {
        return  Array.from(this.articles.values()).flat();
    }

    public getCategories(): Category[] {
        return this.categories
    }

    public addCategory(category: Category) {
        this.categories.push(category)
    }

    public addArticle(categoryPrettyId: string, article: Article) {
        this.articles.get(categoryPrettyId)?.push(article)
    }

    public getArticle(articlePrettyId: string): Article {
        const match = Array.from(this.articles.values()).flat().filter(item => item.prettyId === articlePrettyId)

        if (match.length == 0) {
            throw "failed to find article with id: " + articlePrettyId
        }

        return match[0]
    }
}

export interface DedicatedPage {
    id: string
    prettyId: string
    title: string
    content: string
}

export interface Article {
    id: string
    prettyId: string
    title: string
    description: string
    content: string
}

export interface Category {
    id: string
    prettyId: string
    title: string
}