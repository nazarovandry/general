using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Ball : MonoBehaviour
{
	private float speed = 25f;
	private float jump = 100f;
	private float jumpResistance = 0f;
	private bool canJump = false;
	private float force = 0.4f;
	
	private CircleCollider2D co;
	private Rigidbody2D rb, rb1, rb2, rb3;
	private Transform tr;
	private SpriteRenderer sprite;
	
	private void Awake()
	{
		co = GetComponent<CircleCollider2D>();
		rb = GetComponent<Rigidbody2D>();
		tr = GetComponent<Transform>();
		sprite = GetComponentInChildren<SpriteRenderer>();
		rb1 = GameObject.Find("Circle1").GetComponent<Rigidbody2D>();
		rb2 = GameObject.Find("Circle2").GetComponent<Rigidbody2D>();
		rb3 = GameObject.Find("Circle3").GetComponent<Rigidbody2D>();
	}
	
	//CircleCast(Vector2 origin, float radius, 
	
	void OnCollisionEnter2D(Collision2D c)
	{
		jumpResistance = 0f;
	if (c.collider.IsTouching(co) && c.gameObject.tag != "Player")
		{
			float y = c.contacts[0].normal.y;
			if (y > 0)
			{
				jumpResistance += c.contacts[0].normal.y;
			}
			canJump = true;
		}
	}
	
	void OnCollisionExit2D(Collision2D c)
	{
		canJump = false;
	}
	
	void Update()
    {
		rb.AddForce((GameObject.Find("Circle1").transform.position - tr.position) * force);
		rb.AddForce((GameObject.Find("Circle2").transform.position - tr.position) * force);
		rb.AddForce((GameObject.Find("Circle3").transform.position - tr.position) * force);
		if (Input.GetKey("right") && rb.velocity.x <= 5)
			Right();
		if (Input.GetKey("left") && rb.velocity.x >= -5)
			Left();
		if (Input.GetKey("up") && canJump)
			Up();
    }
	
    private void Right()
    {
		rb.velocity += new Vector2(speed * Time.deltaTime, 0);
	}
	
	private void Left()
    {
		rb.velocity -= new Vector2(speed * Time.deltaTime, 0);
	}
	
	private void Up()
    {
		//rb.velocity += new Vector2(0, jump * jumpResistance * Time.deltaTime);
		rb1.velocity += new Vector2(0, jump * jumpResistance * Time.deltaTime);
		rb2.velocity += new Vector2(0, jump * jumpResistance * Time.deltaTime);
		rb3.velocity += new Vector2(0, jump * jumpResistance * Time.deltaTime);
	}
}
