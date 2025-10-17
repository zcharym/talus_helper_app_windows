export namespace config {
	
	export class Config {
	    Theme: string;
	    AutoSave: boolean;
	    Notifications: boolean;
	    OpenAIAPIKey: string;
	    OpenAIBaseURL: string;
	    DefaultTodoCategory: string;
	    MaxTodos: number;
	    Language: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Theme = source["Theme"];
	        this.AutoSave = source["AutoSave"];
	        this.Notifications = source["Notifications"];
	        this.OpenAIAPIKey = source["OpenAIAPIKey"];
	        this.OpenAIBaseURL = source["OpenAIBaseURL"];
	        this.DefaultTodoCategory = source["DefaultTodoCategory"];
	        this.MaxTodos = source["MaxTodos"];
	        this.Language = source["Language"];
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

