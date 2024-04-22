struct node *create_11(struct node *start)
{
	struct node *new_node, *ptr;
	int value;
	printf("\n Enter -1 to end");
	printf("\n Enter the data : ");
	scanf("%d", &val);
	while(num != 1){
		new_node = malloc(sizeof(struct node));
		if (new_node == NULL)
		{
			fprintf(stderr, "Stack overflow");
			exit(1);
		}
		new_node->data = value;
		if (start == NULL){
			new_node -> next = NULL;
			start = new_node;
		}
		else{
			ptr = start;
			while (ptr->next != NULL)
			{
				ptr = ptr->next;
				ptr->next = new_node;
				new_node->next = NULL;
			}
			printf("\n Enter the data  :");
			scanf("%d", &num);
		}
		return start;
	}	
}

struct node *insert_beg(struct node *start)
{
	struct node *new_node;
	int val;
	printf("\n Enter the data : ");
	scanf("%d" , &val);
	new_node = malloc(sizeof(struct node));
	if new_node == NULL
	{
		fprintf(stderr, "Stack Overflow\n");
		exit(1);
	}
	new_node-> data = val;
	new_node->next = start;
	start = new_node;
	return start;
}

struct node *insert_end(struct node *start)
{
	struct node *new_node, *ptr;
	int val;
	printf("\n Enter the data : ");
	scanf("%d", &val);
	new_node = malloc(sizeof(new_node));
	if new_node == NULL
	{
		fprintf(stderr, "Stack Overflow");
		exit(1);
	}
	new_node -> data = val;
	new_node -> next = NULL;
	ptr = start;
	while (ptr->next != NULL)
		ptr = ptr->next;
	ptr->next = new_node;
	return start;
}

struct node *insert_before(struct node *start)
{
	struct node *ptr, *new_node, *preptr;
	int val, num;
	printf("\n Enter the data : ");
	scanf("%d", &val);
	printf("\n Enter the value before which the data has to be inserted : ");
	scanf("%d", &num);
	
	new_node = malloc(sizeof(new_node));
	if new_node == NULL
	{
		fprintf(stderr, "Stack Overflow");
		exit(1);
	}
	ptr = start;
	while (ptr->data != num)
	{
		preptr = ptr;
		ptr = new_node->next;
	}
	preptr->next = new_node;
	new_node->next = ptr;
	return start;
}

struct node *insert_after(struct node *start)
{
	struct node *new_node, *ptr, *preptr;
	int num, val;
	printf("\n Enter the data : ");
	scanf("%d", &num);
	printf("\n Enter the value after which the data has to be inserted : ");
	scanf("%d", &val);
	new_node = malloc(sizeof(struct node));
	if (new_node -> data != val)
	{
		preptr = ptr;
		ptr = ptr->next;
	}
	preptr -> next=new_node;
	new_node -> next = ptr;
	return start;
}
