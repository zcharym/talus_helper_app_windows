export namespace config {
	
	export class Config {
	    theme: string;
	    autoSave: boolean;
	    notifications: boolean;
	    openAIAPIKey: string;
	    openAIBaseURL: string;
	    defaultTodoCategory: string;
	    maxTodos: number;
	    language: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.theme = source["theme"];
	        this.autoSave = source["autoSave"];
	        this.notifications = source["notifications"];
	        this.openAIAPIKey = source["openAIAPIKey"];
	        this.openAIBaseURL = source["openAIBaseURL"];
	        this.defaultTodoCategory = source["defaultTodoCategory"];
	        this.maxTodos = source["maxTodos"];
	        this.language = source["language"];
	    }
	}

}

export namespace models {
	
	export class Todo {
	    id: string;
	    text: string;
	    completed: boolean;
	    // Go type: time
	    createdAt: any;
	
	    static createFrom(source: any = {}) {
	        return new Todo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.text = source["text"];
	        this.completed = source["completed"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

