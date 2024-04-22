#ifndef __MAIN_H__
#define __MAIN_H__

#include <stdio.h>
#include <stdlib.h>
#include <malloc.h>

struct node
{
	int data;
	struct node *next;
};


struct node *create_11(struct node*);
struct node *display(struct node*);
struct node *insert_beg(struct node *);
struct node *insert_end(struct node *);
struct node *insert_before(struct node *);
struct node *insert_after(struct node*);
struct node *delete_beg(struct node*);
struct node *delete_end(struct node*);
struct node *delete_after(struct node*);
struct node *delete_node(struct node*);
struct node *delete_list(struct node*);
struct node *sort_list(struct node *);

#endif
